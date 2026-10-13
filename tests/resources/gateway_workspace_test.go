package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestGatewayWorkspace(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
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
	})

	t.Run("data source", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("data.konnect_gateway_workspace.by_name", "name", "tf-acceptance-workspace-ds"),
						resource.TestCheckResourceAttr("data.konnect_gateway_workspace.by_name", "description", "A test workspace for the data source"),
						resource.TestCheckResourceAttr("data.konnect_gateway_workspace.by_name", "comment", "A test workspace for the data source"),
						resource.TestCheckResourceAttrSet("data.konnect_gateway_workspace.by_name", "id"),
						resource.TestCheckResourceAttrSet("data.konnect_gateway_workspace.by_name", "created_at"),
						resource.TestCheckResourceAttrPair("data.konnect_gateway_workspace.by_name", "id", "konnect_gateway_workspace.ds", "id"),
						resource.TestCheckOutput("workspace_name", "tf-acceptance-workspace-ds"),
					),
				},
			},
		})
	})

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
						// Every entity below is addressed through the workspace,
						// not the control plane's default one. Plugins are the
						// bulk of the workspace-scoped surface, so both a
						// workspace-global plugin and one attached to an entity
						// inside the workspace are covered.
						resource.TestCheckResourceAttr("konnect_gateway_service.httpbin", "workspace", "tf-acceptance-workspace"),
						resource.TestCheckResourceAttr("konnect_gateway_route.anything", "workspace", "tf-acceptance-workspace"),
						resource.TestCheckResourceAttr("konnect_gateway_consumer.alice", "workspace", "tf-acceptance-workspace"),
						resource.TestCheckResourceAttr("konnect_gateway_key_auth.alice_key", "workspace", "tf-acceptance-workspace"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_rate_limiting.workspace_rl", "workspace", "tf-acceptance-workspace"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_cors.service_cors", "workspace", "tf-acceptance-workspace"),
						// The credential is nested under the consumer, so this
						// also proves the nested workspace-scoped path resolves.
						resource.TestCheckResourceAttrPair(
							"konnect_gateway_key_auth.alice_key", "consumer_id",
							"konnect_gateway_consumer.alice", "id",
						),
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

	// `managed_by` is Optional + Computed. If the API returns anything other
	// than exactly what was configured, Terraform either errors with
	// "provider produced inconsistent result after apply" or plans a diff
	// forever. Isolated in its own subtest so that failure does not mask the
	// workspace-scoped entity coverage above.
	t.Run("managed-by", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "managed_by.terraform", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_workspace.my_gatewayworkspace", "managed_by.team", "platform"),
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

	// Moving an entity between workspaces is the destructive path: `workspace`
	// carries RequiresReplaceIfConfigured because the API cannot relocate an
	// entity. If that plan modifier were ever dropped, this would silently
	// become an in-place update that does nothing.
	t.Run("move-workspace", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// No `workspace` set: the service lands in "default".
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_service.httpbin", "workspace", "default"),
					),
				},
				{
					// Same service, now pinned to the named workspace.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_gateway_service.httpbin", plancheck.ResourceActionReplace),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_service.httpbin", "workspace", "tf-acceptance-workspace-move-target"),
					),
				},
			},
		})
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
						resource.TestCheckResourceAttr("konnect_gateway_config_store_secret.smoke", "key", "smoke-key"),
						resource.TestCheckResourceAttrSet("konnect_gateway_ca_certificate.smoke", "cert_digest"),
						resource.TestCheckResourceAttr("konnect_gateway_custom_plugin_streaming.smoke", "name", "setheader-streaming"),
						resource.TestCheckResourceAttr("konnect_gateway_cloned_plugin.smoke", "name", "smoke-custom-acl"),
						resource.TestCheckResourceAttr("konnect_gateway_cloned_plugin.smoke", "ref", "acl"),
					),
				},
			},
		})
	})
}
