package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestGatewayPluginAiProxy(t *testing.T) {
	t.Parallel()

	resourceName := "konnect_gateway_plugin_ai_proxy.my_plugin"

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
						resource.TestCheckResourceAttr(resourceName, "config.genai_category", "text/generation"),
						resource.TestCheckResourceAttr(resourceName, "config.llm_format", "openai"),
						resource.TestCheckResourceAttr(resourceName, "config.route_type", "llm/v1/chat"),
						resource.TestCheckResourceAttr(resourceName, "config.model.name", "claude-3-5-sonnet-20241022"),
						resource.TestCheckResourceAttr(resourceName, "config.model.provider", "anthropic"),
						resource.TestCheckResourceAttr(resourceName, "config.model.options.max_tokens", "1024"),
						resource.TestCheckResourceAttr(resourceName, "config.model.options.input_cost", "3"),
						resource.TestCheckResourceAttr(resourceName, "config.model.options.output_cost", "15"),
						resource.TestCheckResourceAttr(resourceName, "config.model.options.cache_read_cost", "0.3"),
						resource.TestCheckResourceAttr(resourceName, "config.model.options.cache_write_cost", "3.75"),
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
						resource.TestCheckResourceAttr(resourceName, "config.model.options.max_tokens", "2048"),
						resource.TestCheckResourceAttr(resourceName, "config.model.options.input_cost", "5"),
						resource.TestCheckResourceAttr(resourceName, "config.model.options.output_cost", "20"),
						resource.TestCheckResourceAttr(resourceName, "config.model.options.cache_read_cost", "0.5"),
						resource.TestCheckResourceAttr(resourceName, "config.model.options.cache_write_cost", "4.75"),
					),
				},
			},
		})
	})
}
