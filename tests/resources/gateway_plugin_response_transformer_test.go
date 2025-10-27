package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginResponseTransformer(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_response_transformer.my_response_transformer", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_response_transformer.my_response_transformer", "config.add.headers.0", "New-Header:Header Value"),
					),
				},
				{
					// Update config to use remove instead of add
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_response_transformer.my_response_transformer", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_response_transformer.my_response_transformer", "config.remove.headers.0", "New-Header:Header Value"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_response_transformer.my_response_transformer", "config.add.headers.#", "0"),
					),
				},
			},
		})
	})
}
