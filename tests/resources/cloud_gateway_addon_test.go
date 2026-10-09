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
                tier = "micro"
               }
             }
            }
          }
          owner = {
            control_plane = {
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

		updatedConfig := builder.
			Upsert(cp).
			Upsert(addon.AddAttribute("config.managed_cache.capacity_config.tiered.tier", "small")).
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

						//validating if type is defaults to API.
						resource.TestCheckResourceAttr("konnect_cloud_gateway_addon.my_addon", "owner.control_plane.type", "api"),
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
				// Update step
				{
					Config: updatedConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_cloud_gateway_addon.my_addon",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_cloud_gateway_addon.my_addon",
							"config.managed_cache.capacity_config.tiered.tier",
							"small"),
					),
				},
				{
					Config: updatedConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("Cloud Gateways AddOns owned by AI Gateway", func(t *testing.T) {
		builder := hclbuilder.New()
		aiGateway, err := hclbuilder.FromString(`
          resource "konnect_ai_gateway" "test_ai_gateway" {
             name         = "tf-test-ai-gateway-addon"
             display_name = "TF Test AI Gateway AddOn"
			 deployment_type = "managed"
          }
       `)
		require.NoError(t, err)
		addon, err := hclbuilder.FromString(`
         resource "konnect_cloud_gateway_addon" "my_ai_addon" {
          name     = "tf-test-ai-add-on"

          config = {
            managed_cache = {
             capacity_config = {
               tiered = {
                tier = "micro"
               }
             }
            }
          }
          owner = {
            control_plane = {
             control_plane_geo = "us"
             control_plane_id  = konnect_ai_gateway.test_ai_gateway.id
             type              = "ai"
            }
          }
         }
       `)
		require.NoError(t, err)

		fullConfig := builder.
			Upsert(aiGateway).
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
								"konnect_cloud_gateway_addon.my_ai_addon",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_cloud_gateway_addon.my_ai_addon", "name", "tf-test-ai-add-on"),
						resource.TestCheckResourceAttr("konnect_cloud_gateway_addon.my_ai_addon", "owner.control_plane.type", "ai"),
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
