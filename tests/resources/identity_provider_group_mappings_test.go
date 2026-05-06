package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestIdentityProviderTeamGroupMapping(t *testing.T) {
	t.Run("oidc", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_identity_provider_team_group_mapping.my_mapping", "group", "Tech Leads"),
						resource.TestCheckResourceAttrSet("konnect_identity_provider_team_group_mapping.my_mapping", "id"),
						resource.TestCheckResourceAttrSet("konnect_identity_provider_team_group_mapping.my_mapping", "identity_provider_id"),
						resource.TestCheckResourceAttrSet("konnect_identity_provider_team_group_mapping.my_mapping", "team_id"),
					),
				},
				{
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

	t.Run("saml", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_identity_provider_team_group_mapping.my_mapping", "group", "Engineering"),
						resource.TestCheckResourceAttrSet("konnect_identity_provider_team_group_mapping.my_mapping", "id"),
						resource.TestCheckResourceAttrSet("konnect_identity_provider_team_group_mapping.my_mapping", "identity_provider_id"),
						resource.TestCheckResourceAttrSet("konnect_identity_provider_team_group_mapping.my_mapping", "team_id"),
					),
				},
				{
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
}
