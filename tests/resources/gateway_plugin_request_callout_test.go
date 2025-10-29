package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginRequestCallout(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_request_callout.my_request_callout", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_request_callout.my_request_callout", "config.callouts.0.name", "c1"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_request_callout.my_request_callout", "config.callouts.0.request.url", "http://httpbin.org/uuid"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_request_callout.my_request_callout", "config.callouts.0.request.method", "GET"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_request_callout.my_request_callout", "config.upstream.by_lua"),
					),
				},
				{
					// Update config to change URL
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_request_callout.my_request_callout", "config.callouts.0.request.url", "http://example.com/uuid"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_request_callout.my_request_callout", "config.upstream.by_lua"),
					),
				},
			},
		})
	})
}
