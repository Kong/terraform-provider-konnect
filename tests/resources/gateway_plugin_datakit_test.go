package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginDatakit(t *testing.T) {
	t.Parallel()
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_datakit.my_datakit", "config.nodes.0.static.values", "{\"headers\":{\"Content-Type\":\"application/x-www-form-urlencoded\"}}"),
					),
				},
				{
					// Update some fields
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_datakit.my_datakit", "config.nodes.0.static.values", "{\"headers\":{\"Content-Type\":\"application/json\"}}"),
					),
				},
			},
		})
	})
}
