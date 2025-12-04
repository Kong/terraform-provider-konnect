package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginAiMcpProxy(t *testing.T) {
	t.Parallel()

	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_mcp_proxy.my_plugin", "instance_name", "my-test-plugin"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_mcp_proxy.my_plugin", "config.tools.0.parameters.0.name", "id"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_ai_mcp_proxy.my_plugin", "config.tools.1.request_body.content"),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_mcp_proxy.my_plugin", "config.max_request_body_size", "4096"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_ai_mcp_proxy.my_plugin", "config.tools.0.parameters"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_ai_mcp_proxy.my_plugin", "config.tools.1.request_body"),
					),
				},
			},
		})
	})
}
