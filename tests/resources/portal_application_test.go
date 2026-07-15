package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestPortalApplication(t *testing.T) {
	t.Run("basic-creation", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_portal_application.test_application", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_application.test_application", "name", "test_application_basic"),
						resource.TestCheckResourceAttr("konnect_portal_application.test_application", "owner.type", "developer"),
						resource.TestCheckResourceAttrSet("konnect_portal_application.test_application", "owner.id"),
						resource.TestCheckResourceAttrSet("konnect_portal_application.test_application", "id"),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("auth-strategy-id", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_portal_application.test_application", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_application.test_application", "name", "test_application_auth_id"),
						resource.TestCheckResourceAttr("konnect_portal_application.test_application", "owner.type", "developer"),
						resource.TestCheckResourceAttrSet("konnect_portal_application.test_application", "owner.id"),
						resource.TestCheckResourceAttrSet("konnect_portal_application.test_application", "id"),
						resource.TestCheckResourceAttrSet("konnect_portal_application.test_application", "auth_strategy_id"),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
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
