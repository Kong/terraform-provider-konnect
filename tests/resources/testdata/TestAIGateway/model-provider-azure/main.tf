resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - model-provider-azure"
  name         = "tf-test-aigw-model-provider-azure"
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
