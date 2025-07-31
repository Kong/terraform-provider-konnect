package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestPortalPage(t *testing.T) {
	t.Run("plan-diff", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_page.my_portalpage", "slug", "/my-test-page"),
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

	/*
		todo: fix non-empty refresh plan
			t.Run("update-nullify-fields", func(t *testing.T) {
				resource.Test(t, resource.TestCase{
					ProtoV6ProviderFactories: providerFactory,
					Steps: []resource.TestStep{
						{
							Config:          providerConfigUs,
							ConfigDirectory: config.TestNameDirectory(),
						},
						{
							// Update some fields to null
							Config:          providerConfigUs,
							ConfigDirectory: config.TestStepDirectory(),
							Check: resource.ComposeTestCheckFunc(
								resource.TestCheckNoResourceAttr("konnect_portal_appearance.nullify_appearance_test", "custom_theme"),
								resource.TestCheckNoResourceAttr("konnect_portal_appearance.nullify_appearance_test", "custom_fonts"),
								resource.TestCheckNoResourceAttr("konnect_portal_appearance.nullify_appearance_test", "text"),
								resource.TestCheckNoResourceAttr("konnect_portal_appearance.nullify_appearance_test", "images"),
							),
						},
					},
				})
			})
	*/
}
