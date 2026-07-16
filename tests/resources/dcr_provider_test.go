package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestDcrProvider(t *testing.T) {
	t.Parallel()

	t.Run("kong_identity", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.name", "my-dcr-provider"),
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.issuer", "https://issuer.example.com"),
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.provider_type", "kongIdentity"),
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.labels.team", "platform"),
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "provider_type", "kongIdentity"),
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

	// Regression guard for the perpetual-replace bug tracked as issue 5 in
	// RCA-remaining-issues.md: an application_auth_strategy linked to a
	// dcr_provider via dcr_provider_id must converge to an empty plan, not
	// force replacement on every subsequent plan/apply.
	t.Run("kong_identity linked to auth strategy converges", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
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

	t.Run("requires replacement when issuer changes", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_dcr_provider.my_dcrprovider", plancheck.ResourceActionDestroyBeforeCreate),
						},
					},
				},
			},
		})
	})
}
