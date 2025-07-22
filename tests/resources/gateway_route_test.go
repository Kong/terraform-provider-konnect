package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayRoute(t *testing.T) {
	t.Run("update-nullify-fields", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_route.my_nullify_route", "name", "my-route-name"),
					),
				},
				{
					// Update some fields to null
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_route.my_nullify_route", "name"),
						resource.TestCheckNoResourceAttr("konnect_gateway_route.my_nullify_route", "methods"),
						resource.TestCheckNoResourceAttr("konnect_gateway_route.my_nullify_route", "snis"),
						resource.TestCheckNoResourceAttr("konnect_gateway_route.my_nullify_route", "hosts"),
					),
				},
			},
		})
	})
}
