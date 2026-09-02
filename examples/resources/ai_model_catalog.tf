#example of creating an AI Model in the Konnect Catalog.
resource "konnect_catalog_ai_model" "test_ai_model" {
  name         = "test-ai-model-tf"
  display_name = "Test AI Model"
  description  = "AI Model for testing"
  labels = {
    env = "test"
  }
}

#example of creating a version for the AI Model in the catalog with target models
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

#example of creating a version spec for the AI Model in the catalog with a spec content
resource "konnect_catalog_ai_model_version_spec" "test_ai_model_version_spec" {
  ai_model_id  = konnect_catalog_ai_model.test_ai_model.id
  spec_content = "{\"openapi\":\"3.1.0\",\"info\":{\"title\":\"Test AI Model\",\"version\":\"1.0.0\"},\"paths\":{}}"
  depends_on = [konnect_catalog_ai_model_version.test_ai_model_version]
}


#example of creating an AI Gateway to be used in the AI Model implementation(currently in beta-provider as of 02/09/2026)
resource "konnect_ai_gateway" "test_ai_gateway" {
  provider     = konnect-beta
  name         = "test-ai-gateway-impl"
  display_name = "Test AI Gateway"
}

#example of creating an AI Gateway Model Provider to be used in the AI Model implementation(currently in beta-provider as of 02/09/2026)
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

#example of creating an AI Gateway Model to be used in the AI Model implementation(currently in beta-provider as of 02/09/2026)
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

#example of creating an AI Model implementation in the catalog that uses the AI Gateway and AI Gateway Model
resource "konnect_catalog_ai_model_implementation" "test_ai_model_impl" {
  ai_model_id              = konnect_catalog_ai_model.test_ai_model.id
  gateway_control_plane_id = konnect_ai_gateway.test_ai_gateway.id
  gateway_model_id         = konnect_ai_gateway_model.test_ai_gateway_model.id
}