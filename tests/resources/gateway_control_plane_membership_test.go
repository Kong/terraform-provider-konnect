package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayControlPlaneMembership(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				Config:          providerConfigUs,
				ConfigDirectory: config.TestNameDirectory(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("konnect_gateway_control_plane_membership.group", "control_plane_id"),
					resource.TestCheckResourceAttr("konnect_gateway_control_plane_membership.group", "members.#", "1"),
					resource.TestCheckResourceAttrSet("konnect_gateway_control_plane_membership.group", "members.0.id"),
				),
			},
		},
	})
}
