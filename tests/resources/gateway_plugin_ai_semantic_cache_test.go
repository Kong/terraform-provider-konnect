package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestGatewayPluginAiSemanticCache(t *testing.T) {
	t.Run("without-partial", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// Create: plugin with inline vectordb + embeddings config (no partial)
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "config.embeddings.model.name", "text-embedding-3-small"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "config.embeddings.model.provider", "openai"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "config.vectordb.strategy", "pgvector"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "config.vectordb.dimensions", "1536"),
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

	t.Run("CRUD-with-partial", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// Create: vectordb (pgvector) partial + plugin referencing it
					Config:             providerConfigUs,
					ConfigDirectory:    config.TestNameDirectory(),
					ExpectNonEmptyPlan: true, // Terraform plugin testing throws error on non-refresh plan being non-empty.
					// There is no option to skip that step - so we set our config to pass non-refresh plan empty check.
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_gateway_partial.pgvector_partial", "id"),
						resource.TestCheckResourceAttr("konnect_gateway_partial.pgvector_partial", "vectordb.name", "my_tf_pgvector_partial"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "config.embeddings.model.name", "text-embedding-3-small"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "config.embeddings.model.provider", "openai"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "partials.0.id"),
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "partials.0.path", "config.vectordb"),
					),
				},
				{
					// Update: add a tag (not covered by the partial)
					Config:             providerConfigUs,
					ConfigDirectory:    config.TestStepDirectory(),
					ExpectNonEmptyPlan: true,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "tags.0", "updated"),
						resource.TestCheckResourceAttrSet("konnect_gateway_plugin_ai_semantic_cache.my_ai_semantic_cache", "partials.0.id"),
					),
				},
			},
		})
	})
}
