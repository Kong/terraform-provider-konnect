package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginFileLog(t *testing.T) {
	t.Parallel()

	// We previously broke /dev/stdout support in the file log plugin
	// by enforcing the Lua pattern (converted to regex).
	// This is a smoke test for a very common configuration.
	t.Run("path stdout", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_file_log.test", "config.path", "/dev/stdout"),
					),
				},
			},
		})
	})
}
