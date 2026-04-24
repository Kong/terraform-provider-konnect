package tests

import (
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestGatewayPluginAiCustomGuardrail(t *testing.T) {
	t.Run("create and update", func(t *testing.T) {
		builder := hclbuilder.New()

		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "my_konnect_cp" {
            name = "tf-test-cp-ai-custom-guardrail"
          }
        `)
		require.NoError(t, err)

		aiCustomGuardrail, err := hclbuilder.FromString(`
          resource "konnect_gateway_plugin_ai_custom_guardrail" "my_ai_custom_guardrail" {
            enabled = true

            config = {
              guarding_mode = "INPUT"

              request = {
                url = "https://guardrail.example.com/check"
              }

              response = {
                block         = "blocked"
                block_message = "Request blocked by guardrail"
              }
            }

            control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
          }
        `)
		require.NoError(t, err)

		createConfig := builder.
			Upsert(cp).
			Upsert(aiCustomGuardrail).
			Build()

		updateConfig := builder.
			Upsert(cp).
			Upsert(aiCustomGuardrail.AddAttribute("config.guarding_mode", "BOTH")).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: createConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_plugin_ai_custom_guardrail.my_ai_custom_guardrail",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_custom_guardrail.my_ai_custom_guardrail", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_custom_guardrail.my_ai_custom_guardrail", "config.guarding_mode", "INPUT"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_custom_guardrail.my_ai_custom_guardrail", "config.request.url", "https://guardrail.example.com/check"),
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
								"konnect_gateway_plugin_ai_custom_guardrail.my_ai_custom_guardrail",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_custom_guardrail.my_ai_custom_guardrail", "config.guarding_mode", "BOTH"),
					),
				},
			},
		})
	})
}
