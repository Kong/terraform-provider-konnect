package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestPortalCustomization(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_customization.my_portal_customization", "robots", "my_robots"),
						resource.TestCheckResourceAttr("konnect_portal_customization.my_portal_customization", "theme.name", "my_name"),
						resource.TestCheckResourceAttr("konnect_portal_customization.my_portal_customization", "css", "my_css"),
						resource.TestCheckResourceAttr("konnect_portal_customization.my_portal_customization", "menu.main.0.title", "My Page"),
						resource.TestCheckResourceAttr("konnect_portal_customization.my_portal_customization", "spec_renderer.infinite_scroll", "true"),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_portal_customization.my_portal_customization", "robots"),
						resource.TestCheckNoResourceAttr("konnect_portal_customization.my_portal_customization", "theme"),
						resource.TestCheckNoResourceAttr("konnect_portal_customization.my_portal_customization", "css"),
						resource.TestCheckNoResourceAttr("konnect_portal_customization.my_portal_customization", "menu"),
						resource.TestCheckNoResourceAttr("konnect_portal_customization.my_portal_customization", "spec_renderer"),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PostApplyPostRefresh: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})
}
