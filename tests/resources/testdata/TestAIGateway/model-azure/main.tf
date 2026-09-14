resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - model-azure"
  name         = "tf-test-aigw-model-azure"
}

resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider" {
  provider   = konnect-beta
  gateway_id = konnect_ai_gateway.my_aigateway.id
  azure = {
    config = {
      auth = {
        azure = {
          client_id            = "test-client-id"
          client_secret        = "test-client-id"
          tenant_id            = "test-tenant-id"
          type                 = "azure"
          use_managed_identity = true
        }
      }
      instance = "kong-az-east"
    }
    display_name = "Test TF Azure AI SE"
    name         = "tf-test-azure-ai-provider"
  }
}

resource "konnect_ai_gateway_model" "my_aigatewaymodel_model" {
  provider = konnect-beta
  model = {
    capabilities = [
      "generate"
    ]
    config = {
      route = {
        hosts = []
        model = {
          body_param = "model"
          values = [
            "my-azure-model"
          ]
        }
        paths = [
          "/my-base-path"
        ]
        protocols = [
          "https"
        ]
        https_redirect_status_code = 426
        preserve_host              = false
        regex_priority             = 0
        request_buffering          = true
        response_buffering         = true
        strip_path                 = true
      }
    }
    display_name = "My Test Azure model"
    enabled      = true
    formats = [
      {
        type = "openai"
      }
    ]
    name = "my-azure-model"
    targets = [
      {
        config = {
          azure = {
            api_version   = "2023-05-15"
            deployment_id = "ahagshhh-1sjn-akjnda"
            type          = "azure"
          }
        }
        name     = "gpt-5.6-luna",
        provider = konnect_ai_gateway_model_provider.my_aigatewaymodelprovider.name
      }
    ]
  }
  gateway_id = konnect_ai_gateway.my_aigateway.id
}
