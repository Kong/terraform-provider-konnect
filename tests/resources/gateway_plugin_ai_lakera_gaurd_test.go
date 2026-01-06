package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginAiLakeraGaurd(t *testing.T) {
	t.Parallel()
	t.Run("update-nullify-fields", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_lakera_guard.my_ai_lakera_guard", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_lakera_guard.my_ai_lakera_guard", "config.api_key", "test-api-key-12345"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_lakera_guard.my_ai_lakera_guard", "config.guarding_mode", "INPUT"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_lakera_guard.my_ai_lakera_guard", "config.text_source", "last_message"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_lakera_guard.my_ai_lakera_guard", "config.timeout", "5000"),
					),
				},
				{
					// Update some fields
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_lakera_guard.my_ai_lakera_guard", "enabled", "true"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_ai_lakera_guard.my_ai_lakera_guard", "config.api_key"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_lakera_guard.my_ai_lakera_guard", "config.guarding_mode", "OUTPUT"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_lakera_guard.my_ai_lakera_guard", "config.timeout", "5000"),
					),
				},
			},
		})
	})
}
