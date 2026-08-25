package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestGatewayPluginAiAzureContentSafety(t *testing.T) {
	t.Parallel()

	resourceName := "konnect_gateway_plugin_ai_azure_content_safety.my_plugin"

	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
						resource.TestCheckResourceAttr(resourceName, "config.content_safety_url", "https://my-content-safety.cognitiveservices.azure.com"),
						resource.TestCheckResourceAttr(resourceName, "config.azure_api_version", "2023-10-01"),
						resource.TestCheckResourceAttr(resourceName, "config.guarding_mode", "INPUT"),
						resource.TestCheckResourceAttr(resourceName, "config.output_type", "FourSeverityLevels"),
						resource.TestCheckResourceAttr(resourceName, "config.text_source", "concatenate_all_content"),
						resource.TestCheckResourceAttr(resourceName, "config.response_buffer_size", "100"),
						resource.TestCheckResourceAttr(resourceName, "config.categories.0.name", "Hate"),
						resource.TestCheckResourceAttr(resourceName, "config.categories.0.rejection_level", "4"),
						resource.TestCheckResourceAttr(resourceName, "config.categories.1.name", "Violence"),
						resource.TestCheckResourceAttr(resourceName, "config.categories.1.rejection_level", "2"),
						resource.TestCheckResourceAttr(resourceName, "config.blocklist_names.0", "blocklist-one"),
						resource.TestCheckResourceAttr(resourceName, "config.halt_on_blocklist_hit", "true"),
						resource.TestCheckResourceAttr(resourceName, "config.reveal_failure_reason", "true"),
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
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "config.guarding_mode", "BOTH"),
						resource.TestCheckResourceAttr(resourceName, "config.output_type", "EightSeverityLevels"),
						resource.TestCheckResourceAttr(resourceName, "config.text_source", "concatenate_user_content"),
						resource.TestCheckResourceAttr(resourceName, "config.response_buffer_size", "200"),
						resource.TestCheckResourceAttr(resourceName, "config.categories.0.name", "Hate"),
						resource.TestCheckResourceAttr(resourceName, "config.categories.0.rejection_level", "6"),
						resource.TestCheckResourceAttr(resourceName, "config.categories.1.name", "SelfHarm"),
						resource.TestCheckResourceAttr(resourceName, "config.categories.1.rejection_level", "4"),
						resource.TestCheckResourceAttr(resourceName, "config.blocklist_names.0", "blocklist-two"),
						resource.TestCheckResourceAttr(resourceName, "config.reveal_failure_reason", "false"),
					),
				},
			},
		})
	})
}
