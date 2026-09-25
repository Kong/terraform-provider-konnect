package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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
						resource.TestCheckResourceAttrSet("konnect_identity_provider.oidc_provider", "login_path"),
					),
					// Validate both list (by type filter) and singular (by ID)  identity provider data sources.
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownOutputValueAtPath(
							"identity_provider_list",
							tfjsonpath.New("data").AtSliceIndex(0).AtMapKey("type"),
							knownvalue.StringExact("oidc"),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"identity_provider_list",
							tfjsonpath.New("data").AtSliceIndex(0).AtMapKey("login_path"),
							knownvalue.StringExact("testoidcmapping"),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"identity_provider",
							tfjsonpath.New("type"),
							knownvalue.StringExact("oidc"),
						),
					},
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
						resource.TestCheckResourceAttrSet("konnect_identity_provider.saml_provider", "login_path"),
					),
					// Validate both list and singular identity provider data sources.
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownOutputValueAtPath(
							"identity_provider_list",
							tfjsonpath.New("data").AtSliceIndex(0).AtMapKey("type"),
							knownvalue.StringExact("saml"),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"identity_provider_list",
							tfjsonpath.New("data").AtSliceIndex(0).AtMapKey("login_path"),
							knownvalue.StringExact("testsamlmapping"),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"identity_provider",
							tfjsonpath.New("type"),
							knownvalue.StringExact("saml"),
						),
					},
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
