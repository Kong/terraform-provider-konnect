package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginRateLimitingAdvanced(t *testing.T) {
	t.Run("update-nullify-fields", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "config.header_name", "X-RateLimit-Limit"),
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

	t.Run("use-vault-reference-in-port", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "config.header_name", "X-RateLimit-Limit"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "config.redis.port", "{vault://konnect/redis/port}"),
					),
				},
				{
					// Update redis port to integer
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "config.header_name"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.my_rate_limiting_advanced", "config.redis.port", "6379"),
					),
				},
			},
		})
	})
}
