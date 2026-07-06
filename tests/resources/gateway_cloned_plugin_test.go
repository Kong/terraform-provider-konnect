package tests

import (
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestGatewayClonedPlugin(t *testing.T) {
	t.Run("create and update", func(t *testing.T) {
		builder := hclbuilder.New()

		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "my_konnect_cp" {
            name = "tf-test-cp-cloned-plugin"
          }
        `)
		require.NoError(t, err)

		clonedPlugin, err := hclbuilder.FromString(`
          resource "konnect_gateway_cloned_plugin" "my_cloned_plugin" {
            name = "custom-acl"
            ref  = "acl"

            control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
          }
        `)
		require.NoError(t, err)

		// An instance of the cloned plugin, referencing it by name.
		customACL, err := hclbuilder.FromString(`
          resource "konnect_gateway_custom_plugin" "my_custom_acl" {
            name = "custom-acl"

            config = {
              allow = ["mygroup"]
            }

            control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
          }
        `)
		require.NoError(t, err)
		customACL.DependsOn(clonedPlugin)

		createConfig := builder.
			Upsert(cp).
			Upsert(clonedPlugin).
			Upsert(customACL).
			Build()

		updateConfig := builder.
			Upsert(cp).
			Upsert(clonedPlugin.AddAttribute("priority", "1370058490")).
			Upsert(customACL).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{ // Step 1: create
					Config: createConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_cloned_plugin.my_cloned_plugin",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectResourceAction(
								"konnect_gateway_custom_plugin.my_custom_acl",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_cloned_plugin.my_cloned_plugin", "name", "custom-acl"),
						resource.TestCheckResourceAttr("konnect_gateway_cloned_plugin.my_cloned_plugin", "ref", "acl"),
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin.my_custom_acl", "name", "custom-acl"),
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin.my_custom_acl", "config.allow.0", "mygroup"),
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
								"konnect_gateway_cloned_plugin.my_cloned_plugin",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_cloned_plugin.my_cloned_plugin", "priority", "1370058490"),
					),
				},
			},
		})
	})
}
