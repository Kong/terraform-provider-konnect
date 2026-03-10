package tests

import (
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestCloudGatewayAddOn(t *testing.T) {
	t.Run("Cloud Gateways AddOns", func(t *testing.T) {
		builder := hclbuilder.New()
		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "test_cp" {
             name         = "tf-test-cp-us-external"
             cloud_gateway = true
          }
       `)
		require.NoError(t, err)
		addon, err := hclbuilder.FromString(`
         resource "konnect_cloud_gateway_addon" "my_addon" {
          name     = "tf-test-add-on"

          config = {
            managed_cache = {
             capacity_config = {
               tiered = {
                tier = "small"
               }
             }
            }
          }
          owner = {
            control_plane = {
             kind             = "control-plane"
             control_plane_geo = "us"
             control_plane_id  = konnect_gateway_control_plane.test_cp.id
            }
          }
         }
       `)
		require.NoError(t, err)

		fullConfig := builder.
			Upsert(cp).
			Upsert(addon).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: fullConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_cloud_gateway_addon.my_addon",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_cloud_gateway_addon.my_addon", "name", "tf-test-add-on"),
					),
				},
				{
					Config: fullConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})
}
