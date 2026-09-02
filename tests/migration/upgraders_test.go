// Package migration covers the v0 -> v1 `workspace` state migration.
//
// This file is the cheap tier: it drives every registered upgrader in-process
// with no Konnect credentials and no Terraform CLI, so it runs on every PR. It
// exists because the upgraders are wired up by hand (one wrapper per entity in
// internal/stateupgraders/workspace_upgraders.go), and a resource that gains
// `workspace` but never gets its wrapper fails silently until a user upgrades.
//
// The expensive tier - real provider binaries, real state file, real plan - is
// migration_test.go.
package migration

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	fwschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/kong/terraform-provider-konnect/v3/internal/provider"
)

const (
	workspaceAttr    = "workspace"
	defaultWorkspace = "default"
)

// resourceUnderTest is one row of the migration matrix.
type resourceUnderTest struct {
	typeName     string
	schema       fwschema.Schema
	schemaType   tftypes.Type
	hasWorkspace bool
	upgraders    map[int64]resource.StateUpgrader
}

// discoverResources walks every resource the provider exposes. Importing
// internal/provider also runs its init(), which is what registers each schema
// type with the stateupgraders package - so this mirrors production wiring
// rather than reconstructing it.
func discoverResources(t *testing.T) []resourceUnderTest {
	t.Helper()

	ctx := context.Background()
	p := provider.New("999.99.9")()

	var out []resourceUnderTest
	for _, newResource := range p.Resources(ctx) {
		r := newResource()

		var metaResp resource.MetadataResponse
		r.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "konnect"}, &metaResp)

		var schemaResp resource.SchemaResponse
		r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)

		row := resourceUnderTest{
			typeName:   metaResp.TypeName,
			schema:     schemaResp.Schema,
			schemaType: schemaResp.Schema.Type().TerraformType(ctx),
		}
		_, row.hasWorkspace = schemaResp.Schema.Attributes[workspaceAttr]
		if u, ok := r.(resource.ResourceWithUpgradeState); ok {
			row.upgraders = u.UpgradeState(ctx)
		}

		out = append(out, row)
	}

	if len(out) == 0 {
		t.Fatal("provider exposed no resources")
	}
	return out
}

// TestWorkspaceResourcesAreWiredUp is the guard against the hand-maintained
// wrapper list drifting from the generated schemas: a resource that has a
// `workspace` attribute must also bump its schema version and register a v0
// upgrader, and nothing else should have one.
func TestWorkspaceResourcesAreWiredUp(t *testing.T) {
	t.Parallel()

	var withWorkspace int
	for _, r := range discoverResources(t) {
		t.Run(r.typeName, func(t *testing.T) {
			if !r.hasWorkspace {
				if len(r.upgraders) > 0 {
					t.Errorf("has a state upgrader but no %q attribute - stale wrapper?", workspaceAttr)
				}
				return
			}

			if r.schema.Version != 1 {
				t.Errorf("schema version = %d, want 1 (state written by the previous provider is version 0)", r.schema.Version)
			}
			if _, ok := r.upgraders[0]; !ok {
				t.Errorf("no upgrader registered for schema version 0; add a wrapper in internal/stateupgraders/workspace_upgraders.go")
			}
		})
		if r.hasWorkspace {
			withWorkspace++
		}
	}

	// A generation slip that drops `workspace` everywhere would otherwise make
	// every subtest above vacuously pass.
	if withWorkspace == 0 {
		t.Fatalf("no resource declares a %q attribute", workspaceAttr)
	}
	t.Logf("%d resources carry %q", withWorkspace, workspaceAttr)
}

// TestUpgradeBackfillsWorkspace runs each v0 upgrader against synthetic prior
// state shaped like the current schema minus `workspace` - which is exactly
// what the previous provider version wrote.
func TestUpgradeBackfillsWorkspace(t *testing.T) {
	t.Parallel()

	for _, r := range discoverResources(t) {
		if !r.hasWorkspace {
			continue
		}
		t.Run(r.typeName, func(t *testing.T) {
			upgrader, ok := r.upgraders[0]
			if !ok {
				t.Skip("no v0 upgrader; covered by TestWorkspaceResourcesAreWiredUp")
			}

			t.Run("backfills-when-absent", func(t *testing.T) {
				prior := priorState(r.schema)
				delete(prior, workspaceAttr)

				got := runUpgrader(t, upgrader, r.schemaType, prior)
				if got != defaultWorkspace {
					t.Errorf("workspace = %q, want %q", got, defaultWorkspace)
				}
			})

			// State that already names a workspace must survive untouched, or
			// re-running the upgrade would silently move resources.
			t.Run("preserves-when-present", func(t *testing.T) {
				prior := priorState(r.schema)
				prior[workspaceAttr] = "team-payments"

				got := runUpgrader(t, upgrader, r.schemaType, prior)
				if got != "team-payments" {
					t.Errorf("workspace = %q, want %q", got, "team-payments")
				}
			})
		})
	}
}

// workspaceScopedEntityTypes are the entity types exercised by
// test-terraform/workspace-resources and test-terraform/gateway-workspace:
// resources that live inside a control plane and can be pinned to a named
// workspace instead of "default". Named explicitly (rather than only relying
// on the generic sweep above) so a resource silently dropping its `workspace`
// attribute during regeneration fails here with the fixture it breaks, not
// just an anonymous count.
var workspaceScopedEntityTypes = []string{
	"konnect_gateway_service",
	"konnect_gateway_route",
	"konnect_gateway_route_expression",
	"konnect_gateway_consumer",
	"konnect_gateway_acl",
	"konnect_gateway_key_auth",
	"konnect_gateway_consumer_group",
	"konnect_gateway_vault",
	"konnect_gateway_plugin_cors",
}

// TestWorkspaceScopedEntitiesPreserveNamedWorkspace runs the v0 upgrader for
// every entity type in workspaceScopedEntityTypes against state that already
// names a non-default workspace, mirroring workspace_scoped.tf: a config
// written against a named workspace must come back as that same workspace,
// not get silently reset to "default".
func TestWorkspaceScopedEntitiesPreserveNamedWorkspace(t *testing.T) {
	t.Parallel()

	byType := make(map[string]resourceUnderTest)
	for _, r := range discoverResources(t) {
		byType[r.typeName] = r
	}

	const namedWorkspace = "team-payments"
	for _, typeName := range workspaceScopedEntityTypes {
		t.Run(typeName, func(t *testing.T) {
			r, ok := byType[typeName]
			if !ok {
				t.Fatalf("resource %q not found - test-terraform/workspace-resources fixture is stale", typeName)
			}
			if !r.hasWorkspace {
				t.Fatalf("resource %q has no %q attribute - it can no longer be workspace-scoped", typeName, workspaceAttr)
			}
			upgrader, ok := r.upgraders[0]
			if !ok {
				t.Fatalf("no v0 upgrader registered for %q", typeName)
			}

			prior := priorState(r.schema)
			prior[workspaceAttr] = namedWorkspace

			got := runUpgrader(t, upgrader, r.schemaType, prior)
			if got != namedWorkspace {
				t.Errorf("workspace = %q, want %q", got, namedWorkspace)
			}
		})
	}
}

// TestGatewayWorkspaceResourceIsNotSelfScoped guards the newly added
// konnect_gateway_workspace resource (test-terraform/gateway-workspace):
// a workspace is addressed by (control_plane_id, name), not by a `workspace`
// attribute of its own, so it must never gain one or a state upgrader.
func TestGatewayWorkspaceResourceIsNotSelfScoped(t *testing.T) {
	t.Parallel()

	const typeName = "konnect_gateway_workspace"
	for _, r := range discoverResources(t) {
		if r.typeName != typeName {
			continue
		}
		if r.hasWorkspace {
			t.Errorf("%q has a %q attribute - it should be scoped only by control_plane_id and name", typeName, workspaceAttr)
		}
		if len(r.upgraders) > 0 {
			t.Errorf("%q has a state upgrader but is not workspace-scoped - stale wrapper?", typeName)
		}
		return
	}
	t.Fatalf("resource %q not found - has it been renamed?", typeName)
}

// priorState builds the JSON shape the previous provider wrote: every
// top-level attribute present and null. Nulls are enough - the upgrader only
// reads and writes `workspace`, and everything else just has to survive the
// JSON round-trip and decode against the current schema.
func priorState(s fwschema.Schema) map[string]any {
	raw := make(map[string]any, len(s.Attributes)+len(s.Blocks))
	for name := range s.Attributes {
		raw[name] = nil
	}
	for name := range s.Blocks {
		raw[name] = nil
	}
	return raw
}

// runUpgrader invokes the upgrader and returns the `workspace` value from the
// upgraded state, failing the test on any diagnostic.
func runUpgrader(t *testing.T, upgrader resource.StateUpgrader, schemaType tftypes.Type, prior map[string]any) string {
	t.Helper()

	priorJSON, err := json.Marshal(prior)
	if err != nil {
		t.Fatalf("marshal prior state: %v", err)
	}

	req := resource.UpgradeStateRequest{RawState: &tfprotov6.RawState{JSON: priorJSON}}
	var resp resource.UpgradeStateResponse
	upgrader.StateUpgrader(context.Background(), req, &resp)

	if resp.Diagnostics.HasError() {
		for _, d := range resp.Diagnostics.Errors() {
			t.Errorf("upgrade diagnostic: %s: %s", d.Summary(), d.Detail())
		}
		t.FailNow()
	}
	if resp.DynamicValue == nil {
		t.Fatal("upgrader returned no state")
	}

	upgraded, err := resp.DynamicValue.Unmarshal(schemaType)
	if err != nil {
		t.Fatalf("unmarshal upgraded state: %v", err)
	}

	var attrs map[string]tftypes.Value
	if err := upgraded.As(&attrs); err != nil {
		t.Fatalf("upgraded state is not an object: %v", err)
	}

	var workspace string
	if err := attrs[workspaceAttr].As(&workspace); err != nil {
		t.Fatalf("read %q from upgraded state: %v", workspaceAttr, err)
	}
	return workspace
}
