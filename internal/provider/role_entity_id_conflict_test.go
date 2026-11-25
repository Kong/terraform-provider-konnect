package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The generated entity_id validator comes from x-speakeasy-conflicts-with. This
// runs it against the real resource schemas, so it fails if the annotation is
// dropped or the generator stops emitting it.
func TestRoleEntityIDConflictsWithEntityIDs(t *testing.T) {
	ctx := context.Background()

	resources := map[string]func() resource.Resource{
		"team_role":           NewTeamRoleResource,
		"user_role":           NewUserRoleResource,
		"system_account_role": NewSystemAccountRoleResource,
	}

	known := []string{testControlPlaneID, testWorkspaceID}

	tests := []struct {
		name       string
		entityID   *string
		entityIDs  any // nil: null, tftypes.UnknownValue: unknown, []string: known list
		wantErrors int
	}{
		{name: "entity_id only", entityID: new(testControlPlaneID)},
		{name: "entity_ids only", entityIDs: known},
		{name: "neither"},
		{name: "both", entityID: new(testControlPlaneID), entityIDs: known, wantErrors: 1},
		{name: "entity_ids unknown", entityID: new(testControlPlaneID), entityIDs: tftypes.UnknownValue},
	}

	for rname, newResource := range resources {
		var schemaResp resource.SchemaResponse
		newResource().Schema(ctx, resource.SchemaRequest{}, &schemaResp)
		s := schemaResp.Schema

		attr, ok := s.Attributes["entity_id"].(schema.StringAttribute)
		require.True(t, ok, "%s: entity_id must be a string attribute", rname)
		require.NotEmpty(t, attr.Validators, "%s: entity_id has no validators", rname)

		for _, tt := range tests {
			t.Run(rname+"/"+tt.name, func(t *testing.T) {
				req := validator.StringRequest{
					Path:           path.Root("entity_id"),
					PathExpression: path.MatchRoot("entity_id"),
					Config:         roleConfig(t, s, tt.entityID, tt.entityIDs),
					ConfigValue:    types.StringPointerValue(tt.entityID),
				}

				var resp validator.StringResponse
				for _, v := range attr.Validators {
					v.ValidateString(ctx, req, &resp)
				}

				assert.Equal(t, tt.wantErrors, resp.Diagnostics.ErrorsCount(), "%v", resp.Diagnostics)
			})
		}
	}
}

// roleConfig builds a config for the full resource schema with every attribute
// null except entity_id and entity_ids.
func roleConfig(t *testing.T, s schema.Schema, entityID *string, entityIDs any) tfsdk.Config {
	t.Helper()

	objType, ok := s.Type().TerraformType(context.Background()).(tftypes.Object)
	require.True(t, ok)

	values := make(map[string]tftypes.Value, len(objType.AttributeTypes))
	for name, typ := range objType.AttributeTypes {
		values[name] = tftypes.NewValue(typ, nil)
	}

	if entityID != nil {
		values["entity_id"] = tftypes.NewValue(tftypes.String, *entityID)
	}

	listType := tftypes.List{ElementType: tftypes.String}
	switch v := entityIDs.(type) {
	case nil:
	case []string:
		elems := make([]tftypes.Value, 0, len(v))
		for _, id := range v {
			elems = append(elems, tftypes.NewValue(tftypes.String, id))
		}
		values["entity_ids"] = tftypes.NewValue(listType, elems)
	default:
		values["entity_ids"] = tftypes.NewValue(listType, v)
	}

	return tfsdk.Config{Schema: s, Raw: tftypes.NewValue(objType, values)}
}
