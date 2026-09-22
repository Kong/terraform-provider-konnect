package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginEntitlementEnforcement(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_entitlement_enforcement.my_entitlement_enforcement", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_entitlement_enforcement.my_entitlement_enforcement", "config.entitlement_access_endpoint", "https://entitlements.example.com/v1/access"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_entitlement_enforcement.my_entitlement_enforcement", "config.feature.key", "premium-api-access"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_entitlement_enforcement.my_entitlement_enforcement", "config.customer.field", "x-customer-id"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_entitlement_enforcement.my_entitlement_enforcement", "config.customer.look_up_value_in", "header"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_entitlement_enforcement.my_entitlement_enforcement", "config.response_codes.no_credit_available.http_status", "402"),
					),
				},
				{
					// Update the entitlement access endpoint and a response_codes message.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_entitlement_enforcement.my_entitlement_enforcement", "config.entitlement_access_endpoint", "https://entitlements.example.com/v2/access"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_entitlement_enforcement.my_entitlement_enforcement", "config.response_codes.no_credit_available.message", "Customer has no credit remaining."),
					),
				},
				{
					// Switch look_up_value_in to consumer and remove the field override.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_entitlement_enforcement.my_entitlement_enforcement", "config.customer.look_up_value_in", "consumer"),
						resource.TestCheckNoResourceAttr("konnect_gateway_plugin_entitlement_enforcement.my_entitlement_enforcement", "config.customer.field"),
					),
				},
			},
		})
	})
}
