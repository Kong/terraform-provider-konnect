package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestGatewayPluginServiceProtection(t *testing.T) {
	t.Run("CRUD-with-partial", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// Create: redis_ee partial + service protection plugin referencing it
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_gateway_partial.redis_ee_partial", "id"),
						resource.TestCheckResourceAttr("konnect_gateway_partial.redis_ee_partial", "redis_ee.name", "my_tf_svcprotect_redis_ee_partial"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_service_protection.my_service_protection", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_service_protection.my_service_protection", "config.strategy", "redis"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_service_protection.my_service_protection", "config.namespace", "example"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_service_protection.my_service_protection", "config.error_message", "Service temporarily unavailable"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_service_protection.my_service_protection", "partials.0.id"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_service_protection.my_service_protection", "partials.0.path", "config.redis"),
					),
				},
				{
					// Verify no plan diff after create
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					// Update: change error_message (not covered by the partial)
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_service_protection.my_service_protection", "config.error_message", "Rate limit exceeded"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_service_protection.my_service_protection", "partials.0.id"),
					),
				},
			},
		})
	})
}
