package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginKeyAuth(t *testing.T) {
	t.Parallel()
	t.Run("update-nullify-fields", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_key_auth.my_key_auth", "instance_name", "my_key_auth_plugin"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_key_auth.my_key_auth", "enabled", "true"),
						// This property is computed. It should not be set to a default by provider when not given in config.
						resource.TestCheckResourceAttr("konnect_gateway_plugin_key_auth.my_key_auth", "config.identity_realms.0.scope", "cp"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_key_auth.my_key_auth", "config.anonymous", "anon-user"),
					),
				},
				{
					// Update some fields to null
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_key_auth.my_key_auth", "instance_name"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_key_auth.my_key_auth", "config.anonymous"),
					),
				},
			},
		})
	})
}
