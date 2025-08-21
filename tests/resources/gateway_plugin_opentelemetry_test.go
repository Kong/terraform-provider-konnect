package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginOpenTelemetry(t *testing.T) {
	t.Parallel()
	t.Run("update-nullify-fields", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:             providerConfigUs,
					ConfigDirectory:    config.TestNameDirectory(),
					ExpectNonEmptyPlan: true, // todo: remove after false diff is fixed
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_opentelemetry.my_opentelemetry", "enabled", "true"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_opentelemetry.my_opentelemetry", "config.logs_endpoint"),
					),
				},
				{
					// Update some fields to null
					Config:             providerConfigUs,
					ConfigDirectory:    config.TestStepDirectory(),
					ExpectNonEmptyPlan: true, // todo: remove after false diff is fixed
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_opentelemetry.my_opentelemetry", "config.logs_endpoint"),
					),
				},
			},
		})
	})
}
