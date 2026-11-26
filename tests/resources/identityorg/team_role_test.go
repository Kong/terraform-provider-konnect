package identityorg

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

	// entity_ids carries the one control plane id, mirrored back into computed
	// entity_id, so the implicit second plan stays empty.
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
					// entity_id is echoed on read, so the second plan must be empty.
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

	// Workspace scoped: two positional ids, control plane then workspace.
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
						// entity_id mirrors the first id, and is computed.
						resource.TestCheckResourceAttrPair(
							"konnect_team_role.my_teamrole", "entity_id",
							"konnect_gateway_control_plane.my_konnect_cp", "id",
						),
						resource.TestCheckResourceAttr("konnect_team_role.my_teamrole", "role_name", "Workspace Admin"),
					),
				},
				{
					// entity_ids is ordered, so the second plan must be empty.
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

	// Rejected at plan time: both ids set.
	t.Run("mutually-exclusive", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ExpectError:     regexp.MustCompile(`Attribute "entity_ids" cannot be specified when "entity_id" is\s+specified`),
				},
			},
		})
	})
}
