package tests

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestTeamRole(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_team_role.my_teamrole", "entity_id"),
						resource.TestCheckResourceAttr("konnect_team_role.my_teamrole", "role_name", "Viewer"),
						resource.TestCheckResourceAttr("konnect_team_role.my_teamrole_debug_session", "role_name", "Debug Session Creator"),
					),
				},
			},
		})
	})

	// entity_ids carries the same single control plane id as entity_id does in
	// the CRUD test. The API mirrors it back into entity_id, which is computed,
	// so the step's implicit second plan must come back empty.
	t.Run("entity-ids", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_team_role.my_teamrole", "entity_ids.#", "1"),
						resource.TestCheckResourceAttrPair(
							"konnect_team_role.my_teamrole", "entity_ids.0",
							"konnect_gateway_control_plane.my_konnect_cp", "id",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_team_role.my_teamrole", "entity_id",
							"konnect_gateway_control_plane.my_konnect_cp", "id",
						),
						resource.TestCheckResourceAttr("konnect_team_role.my_teamrole", "role_name", "Viewer"),
					),
				},
				{
					// The API echoes entity_id back on every read. Second plan must
					// be empty, otherwise the attribute is not computed.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	// Workspace scoped assignment: two positional ids, parent control plane
	// first, then the workspace.
	t.Run("workspace-role", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_team_role.my_teamrole", "entity_ids.#", "2"),
						resource.TestCheckResourceAttrPair(
							"konnect_team_role.my_teamrole", "entity_ids.0",
							"konnect_gateway_control_plane.my_konnect_cp", "id",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_team_role.my_teamrole", "entity_ids.1",
							"konnect_gateway_workspace.team_payments", "id",
						),
						// The API mirrors the first positional id into entity_id.
						// entity_id is computed, so this does not replan.
						resource.TestCheckResourceAttrPair(
							"konnect_team_role.my_teamrole", "entity_id",
							"konnect_gateway_control_plane.my_konnect_cp", "id",
						),
						resource.TestCheckResourceAttr("konnect_team_role.my_teamrole", "role_name", "Workspace Admin"),
					),
				},
				{
					// entity_ids is ordered and the API echoes both ids back, so a
					// second plan must be empty.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	// A config that sets both ids is rejected at plan time, before any API call.
	t.Run("mutually-exclusive", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ExpectError:     regexp.MustCompile(`The 'entity_id' and 'entity_ids' fields are mutually exclusive`),
				},
			},
		})
	})
}
