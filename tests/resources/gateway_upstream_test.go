package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayUpstream(t *testing.T) {
	t.Run("update-nullify-fields", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_upstream.my_nullable_upstream", "name", "demo-nullable-upstream"),
					),
				},
				{
					// Update some fields to null
					Config:             providerConfigUs,
					ConfigDirectory:    config.TestStepDirectory(),
					ExpectNonEmptyPlan: true, // Todo - needs to be fixed.
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_gateway_upstream.my_nullable_upstream", "healthchecks"),
						resource.TestCheckNoResourceAttr("konnect_gateway_upstream.my_nullable_upstream", "hash_on_header"),
					),
				},
			},
		})
	})
}
