package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayCustomPlugin(t *testing.T) {
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
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin.custom_basic_auth", "instance_name", "custom-plugin-test"),
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin.custom_basic_auth_nested", "instance_name", "custom-nested-plugin-test"),
					),
				},
			},
		})
	})
}
