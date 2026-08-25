package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestGatewayPluginAiAwsGuardrails(t *testing.T) {
	t.Parallel()

	resourceName := "konnect_gateway_plugin_ai_aws_guardrails.my_plugin"

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
						resource.TestCheckResourceAttr(resourceName, "config.aws_region", "us-east-1"),
						resource.TestCheckResourceAttr(resourceName, "config.guardrails_id", "gr1234567890"),
						resource.TestCheckResourceAttr(resourceName, "config.guardrails_version", "DRAFT"),
						resource.TestCheckResourceAttr(resourceName, "config.guarding_mode", "INPUT"),
						resource.TestCheckResourceAttr(resourceName, "config.text_source", "concatenate_all_content"),
						resource.TestCheckResourceAttr(resourceName, "config.response_buffer_size", "100"),
						resource.TestCheckResourceAttr(resourceName, "config.timeout", "10000"),
						resource.TestCheckResourceAttr(resourceName, "config.allow_masking", "false"),
						resource.TestCheckResourceAttr(resourceName, "config.log_blocked_content", "false"),
						resource.TestCheckResourceAttr(resourceName, "config.stop_on_error", "true"),
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
						resource.TestCheckResourceAttr(resourceName, "config.aws_region", "us-west-2"),
						resource.TestCheckResourceAttr(resourceName, "config.guardrails_version", "1"),
						resource.TestCheckResourceAttr(resourceName, "config.guarding_mode", "BOTH"),
						resource.TestCheckResourceAttr(resourceName, "config.text_source", "concatenate_user_content"),
						resource.TestCheckResourceAttr(resourceName, "config.response_buffer_size", "200"),
						resource.TestCheckResourceAttr(resourceName, "config.timeout", "20000"),
						resource.TestCheckResourceAttr(resourceName, "config.allow_masking", "true"),
						resource.TestCheckResourceAttr(resourceName, "config.log_blocked_content", "true"),
						resource.TestCheckResourceAttr(resourceName, "config.stop_on_error", "false"),
					),
				},
			},
		})
	})
}
