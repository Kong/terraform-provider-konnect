package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// This file uses key-auth, but the things it's testing are not specific to key-auth
// It is checking general plugin behavior
func TestGatewayPlugin(t *testing.T) {
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
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_key_auth.my_key_auth", "ordering.after.access.0"),
					),
				},
				{
					// Update some fields to null
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_key_auth.my_key_auth", "ordering"),
					),
				},
			},
		})
	})
}
