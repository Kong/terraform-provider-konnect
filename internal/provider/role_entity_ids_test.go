package provider

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v3/internal/sdk/models/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// assignRoleModel is the part of the three role resource models that maps the
// Terraform plan to the role assignment API payload. The generated
// code is the same shape for team, user and system account roles, and all three
// need the same entity_ids behaviour, so they are tested together.
type assignRoleModel interface {
	ToSharedAssignRole(ctx context.Context) (*shared.AssignRole, diag.Diagnostics)
}

const (
	testControlPlaneID = "5f2d8a1e-0f6b-4b3e-9a4d-1c2b3a4d5e6f"
	testWorkspaceID    = "7c9e4f2a-3b1d-4c5e-8f6a-9b0c1d2e3f4a"
)

// roleModelFixtures builds one model per role resource, holding a workspace
// scoped assignment: two positional ids, parent first. Pass no ids for a
// model that sets neither field.
func roleModelFixtures(entityIDs ...string) map[string]assignRoleModel {
	var ids []types.String
	for _, id := range entityIDs {
		ids = append(ids, types.StringValue(id))
	}

	return map[string]assignRoleModel{
		"team role": &TeamRoleResourceModel{
			EntityIds:      ids,
			EntityRegion:   types.StringValue("us"),
			EntityTypeName: types.StringValue("Control Planes"),
			RoleName:       types.StringValue("Workspace Admin"),
			TeamID:         types.StringValue("a1b2c3d4-0000-4000-8000-000000000001"),
		},
		"user role": &UserRoleResourceModel{
			EntityIds:      ids,
			EntityRegion:   types.StringValue("us"),
			EntityTypeName: types.StringValue("Control Planes"),
			RoleName:       types.StringValue("Workspace Admin"),
			UserID:         types.StringValue("a1b2c3d4-0000-4000-8000-000000000002"),
		},
		"system account role": &SystemAccountRoleResourceModel{
			AccountID:      types.StringValue("a1b2c3d4-0000-4000-8000-000000000003"),
			EntityIds:      ids,
			EntityRegion:   types.StringValue("us"),
			EntityTypeName: types.StringValue("Control Planes"),
			RoleName:       types.StringValue("Workspace Admin"),
		},
	}
}

// TestRoleEntityIdsRequest asserts the Create request carries entity_ids as an
// ordered list, parent before child, and omits entity_id entirely rather than
// sending it as null.
func TestRoleEntityIdsRequest(t *testing.T) {
	ctx := context.Background()

	for name, model := range roleModelFixtures(testControlPlaneID, testWorkspaceID) {
		t.Run(name, func(t *testing.T) {
			body, diags := model.ToSharedAssignRole(ctx)
			require.False(t, diags.HasError(), "conversion errors: %v", diags)

			require.NotNil(t, body)
			require.Nil(t, body.EntityID, "entity_id must be nil so the payload omits it")
			assert.Equal(t, []string{testControlPlaneID, testWorkspaceID}, body.EntityIds)

			// The wire body is what the API sees, so assert on it directly: the
			// positional order is the whole contract, and a null entity_id would
			// be rejected alongside entity_ids.
			raw, err := json.Marshal(body)
			require.NoError(t, err)

			var payload map[string]any
			require.NoError(t, json.Unmarshal(raw, &payload))

			assert.NotContains(t, payload, "entity_id", "payload: %s", raw)
			assert.Equal(t, []any{testControlPlaneID, testWorkspaceID}, payload["entity_ids"], "payload: %s", raw)
			assert.Equal(t, "Workspace Admin", payload["role_name"])
			assert.Equal(t, "Control Planes", payload["entity_type_name"])
		})
	}
}

// TestRoleEntityIdsOmittedWhenUnset asserts a model that sets neither id sends
// neither key, so the API returns its own 400 rather than reading a null. Before
// x-speakeasy-param-computed the SDK sent "entity_id": null here.
func TestRoleEntityIdsOmittedWhenUnset(t *testing.T) {
	ctx := context.Background()

	for name, model := range roleModelFixtures() {
		t.Run(name, func(t *testing.T) {
			body, diags := model.ToSharedAssignRole(ctx)
			require.False(t, diags.HasError(), "conversion errors: %v", diags)

			require.NotNil(t, body)
			require.Nil(t, body.EntityID)
			assert.Empty(t, body.EntityIds)

			raw, err := json.Marshal(body)
			require.NoError(t, err)

			var payload map[string]any
			require.NoError(t, json.Unmarshal(raw, &payload))

			assert.NotContains(t, payload, "entity_id", "payload: %s", raw)
			assert.NotContains(t, payload, "entity_ids", "payload: %s", raw)
		})
	}
}
