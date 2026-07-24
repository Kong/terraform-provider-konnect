package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestPortalDeveloper(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_portal_developer.test_developer", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_developer.test_developer", "email", "developer@example.com"),
						resource.TestCheckResourceAttr("konnect_portal_developer.test_developer", "full_name", "Test Developer"),
						resource.TestCheckResourceAttr("konnect_portal_developer.test_developer", "status", "approved"),
						resource.TestCheckResourceAttrSet("konnect_portal_developer.test_developer", "id"),
					),
				},

				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_developer.test_developer", "status", "pending"),
					),
				},
			},
		})
	})
}
