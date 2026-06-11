package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestGatewayPluginAiSemanticPromptGuard(t *testing.T) {
	t.Run("without-partial", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// Create: plugin with inline vectordb + embeddings config (no partial)
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "config.embeddings.model.name", "text-embedding-3-small"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "config.embeddings.model.provider", "openai"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "config.vectordb.strategy", "pgvector"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "config.vectordb.dimensions", "1536"),
					),
				},
				{
					// Verify no plan diff after create
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

	t.Run("CRUD-with-embeddings-partial", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// Create: embeddings partial + plugin referencing it; vectordb inline
					Config:             providerConfigUs,
					ConfigDirectory:    config.TestNameDirectory(),
					ExpectNonEmptyPlan: true,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_gateway_partial.embeddings_partial", "id"),
						resource.TestCheckResourceAttr("konnect_gateway_partial.embeddings_partial", "embeddings.name", "my_tf_embeddings_partial"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "config.vectordb.strategy", "pgvector"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "config.vectordb.dimensions", "1536"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "partials.0.id"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "partials.0.path", "config.embeddings"),
					),
				},
				{
					// Update: add a tag (not covered by the partial)
					Config:             providerConfigUs,
					ExpectNonEmptyPlan: true,
					ConfigDirectory:    config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "tags.0", "updated"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_ai_semantic_prompt_guard.my_plugin", "partials.0.id"),
					),
				},
			},
		})
	})
}
