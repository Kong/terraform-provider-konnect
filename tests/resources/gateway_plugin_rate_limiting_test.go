package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginRateLimiting(t *testing.T) {
	t.Run("update-nullify-fields", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting.my_nullable_rate_limiting", "instance_name", "nullable-rl-plugin"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_rate_limiting.my_nullable_rate_limiting", "service.id"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_rate_limiting.my_nullable_rate_limiting", "route.id"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_rate_limiting.my_nullable_rate_limiting", "consumer.id"),
					),
				},
				{
					// Update some fields to null
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_rate_limiting.my_nullable_rate_limiting", "instance_name"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_rate_limiting.my_nullable_rate_limiting", "service"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_rate_limiting.my_nullable_rate_limiting", "route"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_rate_limiting.my_nullable_rate_limiting", "consumer"),
					),
				},
			},
		})
	})

	t.Run("CRUD-with-partial", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// Create the plugin without a partial reference.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting.my_rate_limiting", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting.my_rate_limiting", "config.policy", "redis"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting.my_rate_limiting", "config.hour", "1000"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_rate_limiting.my_rate_limiting", "partials.0.id"),
					),
				},
				{
					// Update the plugin to use the partial.
					Config:             providerConfigUs,
					ConfigDirectory:    config.TestStepDirectory(),
					ExpectNonEmptyPlan: true,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_gateway_partial.redis_ce_partial", "id"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_rate_limiting.my_rate_limiting", "partials.0.id"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting.my_rate_limiting", "partials.0.path", "config.redis"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting.my_rate_limiting", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting.my_rate_limiting", "config.policy", "redis"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting.my_rate_limiting", "config.hour", "1000"),
					),
				},
			},
		})
	})
}
