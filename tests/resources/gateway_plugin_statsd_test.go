package tests

import (
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestGatewayPluginStatsd(t *testing.T) {
	t.Run("create and update", func(t *testing.T) {
		builder := hclbuilder.New()

		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "my_konnect_cp" {
            name = "tf-test-cp-statsd"
          }
        `)
		require.NoError(t, err)

		statsd, err := hclbuilder.FromString(`
          resource "konnect_gateway_plugin_statsd" "my_statsd" {
            enabled = true

            config = {
              host               = "localhost"
              port               = 8125
              allow_status_codes = ["200-205", "400-499"]
              flush_timeout      = 2000
              retry_count        = 10
            }

            control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
          }
        `)
		require.NoError(t, err)

		createConfig := builder.
			Upsert(cp).
			Upsert(statsd).
			Build()

		updateConfig := builder.
			Upsert(cp).
			Upsert(
				statsd.AddAttribute("config.flush_timeout", "5000"),
			).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: createConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_plugin_statsd.my_statsd",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_statsd.my_statsd", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_statsd.my_statsd", "config.host", "localhost"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_statsd.my_statsd", "config.port", "8125"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_statsd.my_statsd", "config.flush_timeout", "2000"),
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
								"konnect_gateway_plugin_statsd.my_statsd",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_statsd.my_statsd", "config.flush_timeout", "5000"),
					),
				},
			},
		})
	})
}
