package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginRateLimitingAdvanced(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "config.header_name", "X-RateLimit-Limit"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "config.redis.port", "6379"),
					),
				},
				{
					// Update some fields to null
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "config.header_name"),
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
					// Create: redis_ee partial + plugin referencing it
					Config:             providerConfigUs,
					ExpectNonEmptyPlan: true,
					ConfigDirectory:    config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_gateway_partial.redis_ee_partial", "id"),
						resource.TestCheckResourceAttr("konnect_gateway_partial.redis_ee_partial", "redis_ee.name", "my_tf_redis_ee_partial"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "config.namespace", "my-namespace"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "partials.0.id"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "partials.0.path", "config.redis"),
					),
				},
				{
					// Update: change config.namespace (not covered by the partial)
					Config:             providerConfigUs,
					ExpectNonEmptyPlan: true,
					ConfigDirectory:    config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "config.namespace", "updated-namespace"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "partials.0.id"),
					),
				},
			},
		})
	})
}
