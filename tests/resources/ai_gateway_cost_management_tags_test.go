package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAIGatewayCostManagementTags(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_cost_management_tags.my_aigatewaycostmanagementtags", "headers.0.tag_key", "tftest-tenant"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_cost_management_tags.my_aigatewaycostmanagementtags", "prefixes.0.name", "tftest-co"),
					),
				},
				{
					// Update headers.0.tag_key
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_cost_management_tags.my_aigatewaycostmanagementtags", "headers.0.tag_key", "tftest-tenant-updated"),
					),
				},
			},
		})
	})
}
