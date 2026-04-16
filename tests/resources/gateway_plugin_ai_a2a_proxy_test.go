package tests

import (
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestGatewayPluginAiA2AProxy(t *testing.T) {
	t.Run("create and update", func(t *testing.T) {
		builder := hclbuilder.New()

		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "my_konnect_cp" {
            name = "tf-test-cp-ai-a2a-proxy"
          }
        `)
		require.NoError(t, err)

		aiA2AProxy, err := hclbuilder.FromString(`
          resource "konnect_gateway_plugin_ai_a2a_proxy" "my_ai_a2a_proxy" {
            enabled = true

            config = {
              max_request_body_size = 1048576
            }

            control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
          }
        `)
		require.NoError(t, err)

		createConfig := builder.
			Upsert(cp).
			Upsert(aiA2AProxy).
			Build()

		updateConfig := builder.
			Upsert(cp).
			Upsert(aiA2AProxy.AddAttribute("config.max_request_body_size", "2097152")).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: createConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_plugin_ai_a2a_proxy.my_ai_a2a_proxy",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_a2a_proxy.my_ai_a2a_proxy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_a2a_proxy.my_ai_a2a_proxy", "config.max_request_body_size", "1048576"),
					),
				},
				{
					Config: createConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: updateConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_plugin_ai_a2a_proxy.my_ai_a2a_proxy",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_a2a_proxy.my_ai_a2a_proxy", "config.max_request_body_size", "2097152"),
					),
				},
			},
		})
	})
}
