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
}
