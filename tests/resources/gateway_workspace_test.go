package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestGatewayWorkspace(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				// Create
				Config:          providerConfigUs,
				ConfigDirectory: config.TestStepDirectory(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "name", "tf-acceptance-workspace"),
					resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "description", "A test workspace for team 1"),
					resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "comment", "A test workspace for team 1"),
					resource.TestCheckResourceAttrSet("konnect_gateway_workspace.my_gatewayworkspace", "created_at"),
					resource.TestCheckResourceAttr("konnect_gateway_service.httpbin", "workspace", "tf-acceptance-workspace"),
				),
			},
			{
				// Re-apply the same config: nothing should have drifted.
				Config:          providerConfigUs,
				ConfigDirectory: config.TestStepDirectory(),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				// Update
				Config:          providerConfigUs,
				ConfigDirectory: config.TestStepDirectory(),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("konnect_gateway_workspace.my_gatewayworkspace", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "description", "A test workspace for team 2"),
					resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "comment", "A test workspace for team 2"),
				),
			},
		},
	})
}

func TestGatewayWorkspaceSmoke(t *testing.T) {
	t.Run("workspaceable-entities", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_service.smoke", "workspace", "tf-acceptance-smoke-workspace"),
						resource.TestCheckResourceAttr("konnect_gateway_route.smoke", "workspace", "tf-acceptance-smoke-workspace"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting_advanced.smoke", "workspace", "tf-acceptance-smoke-workspace"),
					),
				},
			},
		})
	})

	t.Run("non-workspaceable-entities", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_gateway_config_store.smoke", "id"),
						resource.TestCheckResourceAttrSet("konnect_gateway_config_store_secret.smoke", "id"),
					),
				},
			},
		})
	})
}
