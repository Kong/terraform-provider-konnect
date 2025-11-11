package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayConsumerGroup(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_consumer_group.my_gatewayconsumergroup", "name", "TFAcceptanceGWConsumerGroupName"),
						resource.TestCheckResourceAttr("konnect_gateway_consumer_group.my_gatewayconsumergroup", "tags.#", "2"),
						resource.TestCheckResourceAttrSet("konnect_gateway_consumer_group_member.my_gatewayconsumergroupmember", "consumer_group_id"),
					),
				},
				{
					// Update some fields
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_consumer_group.my_gatewayconsumergroup", "tags.#", "1"),
					),
				},
			},
		})
	})
}
