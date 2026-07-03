package tests

import (
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestGatewayCustomPluginStreaming(t *testing.T) {
	t.Run("create and update", func(t *testing.T) {
		builder := hclbuilder.New()

		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "my_konnect_cp" {
            name = "tf-test-cp-custom-plugin-streaming"
          }
        `)
		require.NoError(t, err)

		streamingPlugin, err := hclbuilder.FromString(`
          resource "konnect_gateway_custom_plugin_streaming" "my_streaming_plugin" {
            name = "setheader-streaming"

            handler = <<-EOT
              return {
                VERSION = "1.0.0",
                PRIORITY = 500,
                access = function(self, config)
                  kong.service.request.set_header(config.name, config.value)
                end
              }
            EOT

            schema = <<-EOT
              return {
                name = "setheader-streaming",
                fields = {
                  { protocols = require("kong.db.schema.typedefs").protocols_http },
                  {
                    config = {
                      type = "record",
                      fields = {
                        { name = { description = "The name of the header to set.", type = "string", required = true, }, },
                        { value = { description = "The value for the header.", type = "string", required = true, }, },
                      },
                    },
                  },
                }
              }
            EOT

            control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
          }
        `)
		require.NoError(t, err)

		// An instance of the custom plugin, referencing it by name.
		seth1, err := hclbuilder.FromString(`
          resource "konnect_gateway_custom_plugin" "seth1" {
            name          = "setheader-streaming"
            instance_name = "setheader-instance"

            config = {
              name  = "x-custom-header"
              value = "my-custom-value"
            }

            control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
          }
        `)
		require.NoError(t, err)
		seth1.DependsOn(streamingPlugin)

		createConfig := builder.
			Upsert(cp).
			Upsert(streamingPlugin).
			Upsert(seth1).
			Build()

		updateConfig := builder.
			Upsert(cp).
			Upsert(streamingPlugin.AddAttribute("tags", `["streaming", "set-header"]`)).
			Upsert(seth1).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{ // Step 1: create
					Config: createConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_custom_plugin_streaming.my_streaming_plugin",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectResourceAction(
								"konnect_gateway_custom_plugin.seth1",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin_streaming.my_streaming_plugin", "name", "setheader-streaming"),
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin.seth1", "name", "setheader-streaming"),
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin.seth1", "instance_name", "setheader-instance"),
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin.seth1", "config.name", "x-custom-header"),
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin.seth1", "config.value", "my-custom-value"),
					),
				},
				{ // Step 2: verify no drift
					Config: createConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{ // Step 3: update fields
					Config: updateConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_custom_plugin_streaming.my_streaming_plugin",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin_streaming.my_streaming_plugin", "tags.#", "2"),
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin_streaming.my_streaming_plugin", "tags.0", "streaming"),
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin_streaming.my_streaming_plugin", "tags.1", "set-header"),
					),
				},
			},
		})
	})
}
