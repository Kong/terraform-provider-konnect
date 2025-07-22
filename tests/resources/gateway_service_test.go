package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayService(t *testing.T) {
	t.Run("update-nullify-fields", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_service.httpbin-nullify", "name", "my-service-name"),
					),
				},
				{
					// Update some fields to null
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_service.httpbin-nullify", "name"),
						resource.TestCheckNoResourceAttr("konnect_gateway_service.httpbin-nullify", "path"),
						resource.TestCheckNoResourceAttr("konnect_gateway_service.httpbin-nullify", "ca_certificates"),
						resource.TestCheckNoResourceAttr("konnect_gateway_service.httpbin-nullify", "tls_verify"),
						resource.TestCheckNoResourceAttr("konnect_gateway_service.httpbin-nullify", "tls_verify_depth"),
					),
				},
			},
		})
	})
}
