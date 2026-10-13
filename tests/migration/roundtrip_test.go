package migration

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// The upgrader's contract is narrow: set `workspace`, change nothing else. But
// it implements that by round-tripping the whole state through
// map[string]any - JSON decode, mutate, JSON encode, decode against the
// current schema. Every attribute of every resource rides through that
// round-trip, so a value whose type survives Go's generic JSON decoding badly
// is corrupted for the user with no error.
//
// probeString is deliberately hostile: embedded newlines (PEM bodies,
// lua_schema heredocs), quotes and backslashes (jsonencode'd config), and
// non-ASCII. If JSON escaping is mishandled anywhere, every String attribute
// in the provider trips this at once.
const probeString = "line1\nline2\t\"quoted\" back\\slash ünïcodé 🎉 {\"nested\":\"json\"}"

// populate builds a fully-populated, JSON-encodable value for t: no nulls, so
// every leaf actually exercises the round-trip. Numbers are emitted as
// json.Number so json.Marshal writes the literal digits - that is what lets
// TestUpgradePreservesAllAttributes detect precision loss (see
// TestUpgradeNumberPrecision).
func populate(t tftypes.Type, n json.Number) any {
	switch {
	case t.Is(tftypes.String):
		return probeString
	case t.Is(tftypes.Number):
		return n
	case t.Is(tftypes.Bool):
		return true
	case t.Is(tftypes.List{}):
		return []any{populate(t.(tftypes.List).ElementType, n)}
	case t.Is(tftypes.Set{}):
		return []any{populate(t.(tftypes.Set).ElementType, n)}
	case t.Is(tftypes.Map{}):
		return map[string]any{"probe-key": populate(t.(tftypes.Map).ElementType, n)}
	case t.Is(tftypes.Object{}):
		attrs := t.(tftypes.Object).AttributeTypes
		out := make(map[string]any, len(attrs))
		for name, at := range attrs {
			out[name] = populate(at, n)
		}
		return out
	default:
		// DynamicPseudoType cannot be represented in plain state JSON without
		// its type wrapper; null is the honest value.
		return nil
	}
}

// populatedPriorState is the state the previous provider wrote, with every
// attribute set. Top-level only for `workspace` handling; nested values come
// from populate.
func populatedPriorState(schemaType tftypes.Type, n json.Number) map[string]any {
	root, ok := populate(schemaType, n).(map[string]any)
	if !ok {
		return map[string]any{}
	}
	return root
}

// decodeAgainstSchema decodes raw JSON the same way the upgrader's final step
// does. Used to build the expected value straight from the prior-state JSON,
// bypassing the map[string]any hop - so any corruption introduced by that hop
// shows up as a mismatch.
func decodeAgainstSchema(t *testing.T, schemaType tftypes.Type, raw []byte) tftypes.Value {
	t.Helper()

	v, err := (tfprotov6.RawState{JSON: raw}).UnmarshalWithOpts(schemaType, tfprotov6.UnmarshalOpts{
		ValueFromJSONOpts: tftypes.ValueFromJSONOpts{IgnoreUndefinedAttributes: true},
	})
	if err != nil {
		t.Fatalf("decode reference state: %v", err)
	}
	return v
}

// upgradeRaw runs an upgrader over arbitrary prior-state JSON and returns the
// decoded result. Unlike runUpgrader it does not assume success, so callers
// can assert on diagnostics.
func upgradeRaw(t *testing.T, r resourceUnderTest, priorJSON []byte) (tftypes.Value, error) {
	t.Helper()

	upgrader, ok := r.upgraders[0]
	if !ok {
		t.Fatalf("no v0 upgrader registered for %q", r.typeName)
	}

	req := resource.UpgradeStateRequest{RawState: &tfprotov6.RawState{JSON: priorJSON}}
	var resp resource.UpgradeStateResponse
	upgrader.StateUpgrader(context.Background(), req, &resp)

	if resp.Diagnostics.HasError() {
		var msg string
		for _, d := range resp.Diagnostics.Errors() {
			msg += d.Summary() + ": " + d.Detail() + "; "
		}
		return tftypes.Value{}, fmt.Errorf("%s", msg)
	}
	if resp.DynamicValue == nil {
		return tftypes.Value{}, fmt.Errorf("upgrader returned no state")
	}

	v, err := resp.DynamicValue.Unmarshal(r.schemaType)
	if err != nil {
		return tftypes.Value{}, fmt.Errorf("unmarshal upgraded state: %w", err)
	}
	return v, nil
}

// TestUpgradePreservesAllAttributes is the exhaustive round-trip check: for
// every workspace-carrying resource, feed prior state with every attribute
// populated and assert the upgraded value is byte-for-byte what a direct
// decode of that same state would produce, once `workspace` is accounted for.
//
// This is what the null-filled prior state in upgraders_test.go cannot do: a
// state full of nulls round-trips through map[string]any trivially, so it
// proves nothing about real values.
func TestUpgradePreservesAllAttributes(t *testing.T) {
	t.Parallel()

	// Comfortably inside float64's exact-integer range, so this test reports
	// genuine corruption rather than the known precision boundary that
	// TestUpgradeNumberPrecision pins down separately.
	const safeNumber = json.Number("1756614526")

	for _, r := range discoverResources(t) {
		if !r.hasWorkspace {
			continue
		}
		t.Run(r.typeName, func(t *testing.T) {
			prior := populatedPriorState(r.schemaType, safeNumber)
			delete(prior, workspaceAttr)

			priorJSON, err := json.Marshal(prior)
			if err != nil {
				t.Fatalf("marshal prior state: %v", err)
			}

			got, err := upgradeRaw(t, r, priorJSON)
			if err != nil {
				t.Fatalf("upgrade populated state: %v", err)
			}

			// The expected value is the same state with `workspace` filled in,
			// decoded directly - never through map[string]any.
			prior[workspaceAttr] = defaultWorkspace
			expectedJSON, err := json.Marshal(prior)
			if err != nil {
				t.Fatalf("marshal expected state: %v", err)
			}
			want := decodeAgainstSchema(t, r.schemaType, expectedJSON)

			if !got.Equal(want) {
				t.Errorf("upgraded state differs from a direct decode of the same prior state\n got: %s\nwant: %s", got.String(), want.String())
			}
		})
	}
}

// TestUpgradeWorkspaceEdgeCases covers every shape `workspace` can arrive in.
// Absent is the ordinary pre-migration case; null and empty string are what
// hand-edited or partially-migrated state looks like; an explicit value must
// never be overwritten, or a second upgrade pass would relocate resources.
func TestUpgradeWorkspaceEdgeCases(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		// set mutates prior state to produce the shape under test.
		set  func(m map[string]any)
		want string
	}{
		{"absent", func(m map[string]any) { delete(m, workspaceAttr) }, defaultWorkspace},
		{"null", func(m map[string]any) { m[workspaceAttr] = nil }, defaultWorkspace},
		{"empty-string", func(m map[string]any) { m[workspaceAttr] = "" }, defaultWorkspace},
		{"explicit-default", func(m map[string]any) { m[workspaceAttr] = defaultWorkspace }, defaultWorkspace},
		{"named", func(m map[string]any) { m[workspaceAttr] = "team-payments" }, "team-payments"},
		// Workspace names are user-supplied; they must not be normalised.
		{"named-with-hyphens-and-digits", func(m map[string]any) { m[workspaceAttr] = "team-1_ops.v2" }, "team-1_ops.v2"},
	}

	for _, r := range discoverResources(t) {
		if !r.hasWorkspace {
			continue
		}
		t.Run(r.typeName, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					prior := priorState(r.schema)
					tc.set(prior)

					priorJSON, err := json.Marshal(prior)
					if err != nil {
						t.Fatalf("marshal prior state: %v", err)
					}
					got, err := upgradeRaw(t, r, priorJSON)
					if err != nil {
						t.Fatalf("upgrade: %v", err)
					}

					var attrs map[string]tftypes.Value
					if err := got.As(&attrs); err != nil {
						t.Fatalf("upgraded state is not an object: %v", err)
					}
					var ws string
					if err := attrs[workspaceAttr].As(&ws); err != nil {
						t.Fatalf("read %q: %v", workspaceAttr, err)
					}
					if ws != tc.want {
						t.Errorf("workspace = %q, want %q", ws, tc.want)
					}
				})
			}
		})
	}
}

// TestUpgradeToleratesForeignState covers state the current schema no longer
// describes. Terraform hands the upgrader whatever the old provider wrote, so
// an attribute that has since been renamed or removed must be dropped
// silently - that is what IgnoreUndefinedAttributes is for - rather than
// failing the upgrade and stranding the user on the old provider.
func TestUpgradeToleratesForeignState(t *testing.T) {
	t.Parallel()

	for _, r := range discoverResources(t) {
		if !r.hasWorkspace {
			continue
		}
		t.Run(r.typeName, func(t *testing.T) {
			prior := priorState(r.schema)
			delete(prior, workspaceAttr)
			prior["attribute_removed_in_a_later_release"] = "stale"
			prior["another_stale_attribute"] = map[string]any{"nested": []any{1, 2, 3}}

			priorJSON, err := json.Marshal(prior)
			if err != nil {
				t.Fatalf("marshal prior state: %v", err)
			}

			got, err := upgradeRaw(t, r, priorJSON)
			if err != nil {
				t.Fatalf("upgrade with unknown attributes: %v", err)
			}

			var attrs map[string]tftypes.Value
			if err := got.As(&attrs); err != nil {
				t.Fatalf("upgraded state is not an object: %v", err)
			}
			if _, leaked := attrs["attribute_removed_in_a_later_release"]; leaked {
				t.Error("unknown attribute leaked into upgraded state")
			}
			var ws string
			if err := attrs[workspaceAttr].As(&ws); err != nil {
				t.Fatalf("read %q: %v", workspaceAttr, err)
			}
			if ws != defaultWorkspace {
				t.Errorf("workspace = %q, want %q", ws, defaultWorkspace)
			}
		})
	}
}

// TestUpgradeRejectsMalformedState pins the failure mode for state the
// upgrader cannot parse. Terraform surfaces diagnostics to the user; a panic
// crashes the provider plugin and produces an unactionable error, so the
// distinction matters.
func TestUpgradeRejectsMalformedState(t *testing.T) {
	t.Parallel()

	var sample resourceUnderTest
	for _, r := range discoverResources(t) {
		if r.hasWorkspace {
			if _, ok := r.upgraders[0]; ok {
				sample = r
				break
			}
		}
	}
	if sample.typeName == "" {
		t.Fatal("no workspace-carrying resource with a v0 upgrader")
	}

	for _, tc := range []struct {
		name  string
		state string
	}{
		{"truncated-json", `{"id": `},
		{"not-an-object", `["id"]`},
		{"bare-string", `"nope"`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// A panic here fails the test rather than taking the suite down.
			if _, err := upgradeRaw(t, sample, []byte(tc.state)); err == nil {
				t.Error("expected a diagnostic, got a successful upgrade")
			}
		})
	}

	// An empty object is malformed-adjacent but legitimate: it is what state
	// for a resource with every attribute null looks like. It must succeed.
	t.Run("empty-object", func(t *testing.T) {
		if _, err := upgradeRaw(t, sample, []byte(`{}`)); err != nil {
			t.Errorf("empty prior state should upgrade cleanly, got: %v", err)
		}
	})
}

// TestUpgradeNumberPrecision pins the one lossy step in the round-trip: the
// upgrader decodes prior state into map[string]any, which turns every JSON
// number into a float64. Integers above 2^53 cannot be represented exactly,
// so they would be silently rewritten.
//
// This is currently harmless - Konnect identifiers are UUID strings and the
// numeric attributes in these schemas are ports, counts and Unix timestamps,
// all far below the boundary. The test exists so that if a schema ever gains
// a large integer attribute, this fails loudly instead of corrupting state.
func TestUpgradeNumberPrecision(t *testing.T) {
	t.Parallel()

	var sample resourceUnderTest
	for _, r := range discoverResources(t) {
		if !r.hasWorkspace {
			continue
		}
		if _, ok := r.upgraders[0]; !ok {
			continue
		}
		if numericAttr(r.schemaType) != "" {
			sample = r
			break
		}
	}
	if sample.typeName == "" {
		t.Skip("no workspace-carrying resource with a top-level numeric attribute")
	}
	attr := numericAttr(sample.schemaType)

	// 2^53 is the largest integer float64 represents exactly.
	const withinFloat64 = json.Number("9007199254740992")

	prior := priorState(sample.schema)
	delete(prior, workspaceAttr)
	prior[attr] = withinFloat64

	priorJSON, err := json.Marshal(prior)
	if err != nil {
		t.Fatalf("marshal prior state: %v", err)
	}
	got, err := upgradeRaw(t, sample, priorJSON)
	if err != nil {
		t.Fatalf("upgrade: %v", err)
	}

	prior[workspaceAttr] = defaultWorkspace
	expectedJSON, err := json.Marshal(prior)
	if err != nil {
		t.Fatalf("marshal expected state: %v", err)
	}
	want := decodeAgainstSchema(t, sample.schemaType, expectedJSON)

	if !got.Equal(want) {
		t.Errorf("%s.%s was not preserved across the upgrade\n got: %s\nwant: %s",
			sample.typeName, attr, got.String(), want.String())
	}
}

// numericAttr returns the name of a top-level Number attribute, or "".
func numericAttr(schemaType tftypes.Type) string {
	obj, ok := schemaType.(tftypes.Object)
	if !ok {
		return ""
	}
	for name, at := range obj.AttributeTypes {
		if at.Is(tftypes.Number) {
			return name
		}
	}
	return ""
}
