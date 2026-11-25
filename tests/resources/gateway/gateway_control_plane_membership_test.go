package gateway

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
					resource.TestCheckResourceAttrSet("konnect_gateway_control_plane_membership.group_member", "id"),
					resource.TestCheckResourceAttr("konnect_gateway_control_plane_membership.group_member", "members.#", "1"),
					resource.TestCheckResourceAttr("konnect_gateway_control_plane_membership.group_member2", "members.#", "1"),
					resource.TestCheckResourceAttrSet("konnect_gateway_control_plane_membership.group_member", "members.0.id"),
					resource.TestCheckResourceAttrSet("konnect_gateway_control_plane_membership.group_member2", "members.0.id"),
				),
			},
		},
	})
}
