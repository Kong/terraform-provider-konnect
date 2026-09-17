resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "My AI Gateway"
  name         = "my-ai-gateway"
}

resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  anthropic = {
    name         = "anthropic-provider"
    display_name = "Anthropic AI Provider"
    config = {
      auth = {
        params = [
          {
            location = "body"
            name     = "x-api-key"
            value    = "...my_anthropic_api_key..."
          }
        ]
      }
    }
  }
}

resource "konnect_ai_gateway_model" "my_aigatewaymodel" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  api = {
    capabilities = [
      "batches"
    ]
    name         = "claude-model"
    display_name = "Claude Model"
    enabled      = true
    formats = [
      {
        type = "anthropic"
      }
    ]
    config = {
      route = {
        paths = ["/claude"]
      }
    }
    targets = [
      {
        name     = "claude-5"
        provider = konnect_ai_gateway_model_provider.my_aigatewaymodelprovider.anthropic.name
        config = {
          anthropic = {
            max_tokens    = 1024
            temperature   = 0.7
          }
        }
      }
    ]
  }
}
