package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

const (
	catalogAiModel = `
		resource "konnect_catalog_ai_model" "test_ai_model" {
			name         = "test-ai-model-tf"
			display_name = "Test AI Model"
			description  = "AI Model for testing"
			labels = {
				env = "test"
			}
		}
	`

	testAiGateway = `
		resource "konnect_ai_gateway" "test_ai_gateway" {
			provider     = konnect-beta
			name         = "test-ai-gateway-impl"
			display_name = "Test AI Gateway"
		}
	`

	testAiGatewayModelProvider = `
		resource "konnect_ai_gateway_model_provider" "test_openai_provider" {
			provider   = konnect-beta
			gateway_id = konnect_ai_gateway.test_ai_gateway.id
			openai = {
				name         = "openai"
				display_name = "OpenAI"
				config = {
					auth = {}
				}
			}
		}
	`

	testAiGatewayModel = `
		resource "konnect_ai_gateway_model" "test_ai_gateway_model" {
			provider   = konnect-beta
			gateway_id = konnect_ai_gateway.test_ai_gateway.id
			model = {
				name         = "gpt-4o-model"
				display_name = "Test GPT 4o model"
				capabilities = ["generate"]
				formats = [
					{ type = "openai" }
				]
				config = {
					route = {
						route_type = "direct"
						hosts = []
					}
				}
				targets = [
					{
						name     = "gpt-4o"
						provider = konnect_ai_gateway_model_provider.test_openai_provider.openai.name
						config = {
							openai = {}
						}
					}
				]
			}
		}
	`

	catalogAiModelImplementation = `
		resource "konnect_catalog_ai_model_implementation" "test_ai_model_impl" {
			ai_model_id              = konnect_catalog_ai_model.test_ai_model.id
			gateway_control_plane_id = konnect_ai_gateway.test_ai_gateway.id
			gateway_model_id         = konnect_ai_gateway_model.test_ai_gateway_model.id
		}
	`

	catalogAiModelVersion = `
		resource "konnect_catalog_ai_model_version" "test_ai_model_version" {
			ai_model_id = konnect_catalog_ai_model.test_ai_model.id
			version     = "1.0.0"
			target_models = [
				{
					name     = "gpt-4o"
					provider = "openai"
				},
				{
					name     = "gpt-4o-mini"
					provider = "openai"
				}
			]
		}
	`

	catalogAiModelVersionSpec = `
		resource "konnect_catalog_ai_model_version_spec" "test_ai_model_version_spec" {
			ai_model_id  = konnect_catalog_ai_model.test_ai_model.id
			spec_content = "{\"openapi\":\"3.1.0\",\"info\":{\"title\":\"Test AI Model\",\"version\":\"1.0.0\"},\"paths\":{}}"
			depends_on = [konnect_catalog_ai_model_version.test_ai_model_version]
		}
	`
)

func TestCatalogAiModel(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("Catalog AI Model", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		aiModel, err := hclbuilder.FromString(catalogAiModel)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(aiModel).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_catalog_ai_model.test_ai_model", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_catalog_ai_model.test_ai_model", "name", "test-ai-model-tf"),
						resource.TestCheckResourceAttr("konnect_catalog_ai_model.test_ai_model", "display_name", "Test AI Model"),
					),
				},
				{
					Config: builder.Upsert(aiModel).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(aiModel.AddAttribute("description", `"Updated AI Model for testing"`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_catalog_ai_model.test_ai_model", "description", "Updated AI Model for testing"),
					),
				},
			},
		})
	})

	t.Run("Catalog AI Model Implementation", func(t *testing.T) {
		builder := hclbuilder.New()

		aiModel, err := hclbuilder.FromString(catalogAiModel)
		require.NoError(t, err)

		gateway, err := hclbuilder.FromString(testAiGateway)
		require.NoError(t, err)

		modelProvider, err := hclbuilder.FromString(testAiGatewayModelProvider)
		require.NoError(t, err)

		model, err := hclbuilder.FromString(testAiGatewayModel)
		require.NoError(t, err)

		impl, err := hclbuilder.FromString(catalogAiModelImplementation)
		require.NoError(t, err)

		baseConfig := builder.Upsert(aiModel).Upsert(gateway).Upsert(modelProvider).Upsert(model).Upsert(impl).Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			ExternalProviders: map[string]resource.ExternalProvider{
				"konnect-beta": {
					Source: "Kong/konnect-beta",
				},
			},
			Steps: []resource.TestStep{
				{
					Config: baseConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_catalog_ai_model_implementation.test_ai_model_impl", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_catalog_ai_model_implementation.test_ai_model_impl", "id"),
						resource.TestCheckResourceAttrSet("konnect_catalog_ai_model_implementation.test_ai_model_impl", "ai_model_id"),
					),
				},
			},
		})
	})

	t.Run("Catalog AI Model Version", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		aiModel, err := hclbuilder.FromString(catalogAiModel)
		require.NoError(t, err)

		version, err := hclbuilder.FromString(catalogAiModelVersion)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(aiModel).Upsert(version).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_catalog_ai_model_version.test_ai_model_version", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_catalog_ai_model_version.test_ai_model_version", "id"),
						resource.TestCheckResourceAttr("konnect_catalog_ai_model_version.test_ai_model_version", "version", "1.0.0"),
					),
				},
				{
					Config: builder.Upsert(aiModel).Upsert(version).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(aiModel).Upsert(version.AddAttribute("version", `"2.0.0"`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_catalog_ai_model_version.test_ai_model_version", "version", "2.0.0"),
					),
				},
			},
		})
	})

	t.Run("Catalog AI Model Version Spec", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		aiModel, err := hclbuilder.FromString(catalogAiModel)
		require.NoError(t, err)

		version, err := hclbuilder.FromString(catalogAiModelVersion)
		require.NoError(t, err)

		spec, err := hclbuilder.FromString(catalogAiModelVersionSpec)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(aiModel).Upsert(version).Upsert(spec).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_catalog_ai_model_version_spec.test_ai_model_version_spec", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_catalog_ai_model_version_spec.test_ai_model_version_spec", "ai_model_id"),
						resource.TestCheckResourceAttr("konnect_catalog_ai_model_version_spec.test_ai_model_version_spec", "spec_content", "{\"openapi\":\"3.1.0\",\"info\":{\"title\":\"Test AI Model\",\"version\":\"1.0.0\"},\"paths\":{}}"),
					),
				},
				{
					Config: builder.Upsert(aiModel).Upsert(version).Upsert(spec).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(aiModel).Upsert(version).Upsert(spec.AddAttribute("spec_content", `"{\"openapi\":\"3.1.0\",\"info\":{\"title\":\"Updated AI Model\",\"version\":\"2.0.0\"},\"paths\":{}}"`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_catalog_ai_model_version_spec.test_ai_model_version_spec", "spec_content", "{\"openapi\":\"3.1.0\",\"info\":{\"title\":\"Updated AI Model\",\"version\":\"2.0.0\"},\"paths\":{}}"),
					),
				},
			},
		})
	})
}
