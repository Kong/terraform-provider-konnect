resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - model-provider-anthropic"
  name         = "tf-test-aigw-model-provider-anthropic"
}

resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider_anthropic" {
  provider   = konnect-beta
  gateway_id = konnect_ai_gateway.my_aigateway.id
  anthropic = {
    config = {
      auth = {
        headers = []
        params = [
          {
            location = "body"
            name     = "param_name"
            value    = "...my_value..."
          }
        ]
      }
    }
    display_name = "TF Test Anthropic AI Provider"
    name         = "tf-test-anthropic-provider"
  }
}
