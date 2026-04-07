package tests

import (
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestGatewayPluginKafkaUpstream(t *testing.T) {
	t.Run("create and update", func(t *testing.T) {
		builder := hclbuilder.New()

		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "my_konnect_cp" {
            name = "tf-test-cp-kafka-upstream"
          }
        `)
		require.NoError(t, err)

		kafkaUpstream, err := hclbuilder.FromString(`
          resource "konnect_gateway_plugin_kafka_upstream" "my_kafka_upstream" {
            enabled = true

            config = {
              topic = "my-topic"
            }

            control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
          }
        `)
		require.NoError(t, err)

		createConfig := builder.
			Upsert(cp).
			Upsert(kafkaUpstream).
			Build()

		updateConfig := builder.
			Upsert(cp).
			Upsert(kafkaUpstream.AddAttribute("config.topic", "my-updated-topic")).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: createConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_plugin_kafka_upstream.my_kafka_upstream",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_kafka_upstream.my_kafka_upstream", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_kafka_upstream.my_kafka_upstream", "config.topic", "my-topic"),
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
								"konnect_gateway_plugin_kafka_upstream.my_kafka_upstream",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_kafka_upstream.my_kafka_upstream", "config.topic", "my-updated-topic"),
					),
				},
			},
		})
	})
}
