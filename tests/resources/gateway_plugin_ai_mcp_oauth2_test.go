package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginAiMcpOauth2(t *testing.T) {
	t.Parallel()

	// We create a plugin definition for a built in plugin, but use
	// the custom plugin resource to ensure it works
	t.Run("smoke", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_mcp_oauth2.my_plugin", "instance_name", "my-test-plugin"),
					),
				},
			},
		})
	})
}
