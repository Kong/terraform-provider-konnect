package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestGatewayWorkspace(t *testing.T) {
	t.Run("plan-diff", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "name", "tf-acceptance-workspace"),
						resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "description", "A test workspace for team 1"),
						resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "comment", "A test workspace for team 1"),
						resource.TestCheckResourceAttrSet("konnect_gateway_workspace.my_gatewayworkspace", "created_at"),
						// The service is addressed through the workspace, not the
						// control plane's default one.
						resource.TestCheckResourceAttr("konnect_gateway_service.httpbin", "workspace", "tf-acceptance-workspace"),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("update", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "description", "A test workspace for team 1"),
					),
				},
				{
					// Update some fields
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
	})
}
