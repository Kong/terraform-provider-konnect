package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestOrgPATSettings(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			Steps: []resource.TestStep{
				{
					ProtoV6ProviderFactories: providerFactory,
					Config:                   providerConfigUs,
					ConfigDirectory:          config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_organization_personal_access_token_settings.my_organizationpersonalaccesstokensettings", "pats_enabled", "true"),
						resource.TestCheckResourceAttr("konnect_organization_personal_access_token_settings.my_organizationpersonalaccesstokensettings", "max_expiration_period_days", "364"),
					),
				},
				{
					ProtoV6ProviderFactories: providerFactory,
					Config:                   providerConfigUs,
					ConfigDirectory:          config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_organization_personal_access_token_settings.my_organizationpersonalaccesstokensettings", "max_expiration_period_days", "365"),
					),
				},
				{
					ProtoV6ProviderFactories: providerFactory,
					Config:                   providerConfigUs,
					ConfigDirectory:          config.TestStepDirectory(),
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
