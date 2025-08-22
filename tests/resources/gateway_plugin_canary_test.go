package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginCanary(t *testing.T) {
	t.Parallel()
	t.Run("update-nullify-fields", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_canary.my_canary", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_canary.my_canary", "config.canary_by_header_name", "X-Canary-User"),
					),
				},
				{
					// Update some fields - change strategy
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_canary.my_canary", "config.canary_by_header_name"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_canary.my_canary", "config.steps", "5"),
					),
				},
			},
		})
	})
}
