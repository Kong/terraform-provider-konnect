package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginCors(t *testing.T) {
	t.Parallel()
	t.Run("update-nullify-fields", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_cors.my_cors", "enabled", "true"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_cors.my_cors", "config.origins.0"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_cors.my_cors", "config.headers.0"),
					),
				},
				{
					// Update some fields to null
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_cors.my_cors", "config.origins"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_cors.my_cors", "config.headers"),
					),
				},
			},
		})
	})
}
