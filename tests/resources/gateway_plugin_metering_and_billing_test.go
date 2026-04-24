package tests

import (
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestGatewayPluginMeteringAndBilling(t *testing.T) {
	t.Run("create and update", func(t *testing.T) {
		builder := hclbuilder.New()

		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "my_konnect_cp" {
            name = "tf-test-cp-metering-and-billing"
          }
        `)
		require.NoError(t, err)

		meteringAndBilling, err := hclbuilder.FromString(`
          resource "konnect_gateway_plugin_metering_and_billing" "my_metering_and_billing" {
            enabled = true

            config = {
              api_token       = "my-api-token"
              ingest_endpoint = "https://ingest.example.com/usage"
            }

            control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
          }
        `)
		require.NoError(t, err)

		createConfig := builder.
			Upsert(cp).
			Upsert(meteringAndBilling).
			Build()

		updateConfig := builder.
			Upsert(cp).
			Upsert(meteringAndBilling.AddAttribute("config.ingest_endpoint", "https://ingest-updated.example.com/usage")).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: createConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_plugin_metering_and_billing.my_metering_and_billing",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_metering_and_billing.my_metering_and_billing", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_metering_and_billing.my_metering_and_billing", "config.ingest_endpoint", "https://ingest.example.com/usage"),
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
								"konnect_gateway_plugin_metering_and_billing.my_metering_and_billing",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_metering_and_billing.my_metering_and_billing", "config.ingest_endpoint", "https://ingest-updated.example.com/usage"),
					),
				},
			},
		})
	})
}
