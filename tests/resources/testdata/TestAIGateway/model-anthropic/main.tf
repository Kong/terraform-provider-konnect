resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - model-anthropic"
  name         = "tf-test-aigw-model-anthropic"
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

resource "konnect_ai_gateway_model" "my_aigatewaymodel" {
  provider = konnect-beta
  api = {
    capabilities = [
      "batches"
    ]
    config = {
      route = {
        hosts = []
        paths = [
          "/my-test-path"
        ]
        https_redirect_status_code = 426
        preserve_host              = false
        regex_priority             = 0
        request_buffering          = true
        response_buffering         = true
        strip_path                 = true
      }
    }
    display_name = "My Test Claude 5 model"
    enabled      = true
    formats = [
      {
        type = "anthropic"
      }
    ]
    name = "tf-test-claude-5-model"
    targets = [
      {
        config = {
          anthropic = {
            embeddings_dimensions = 7
            input_cost            = 3.7
            max_tokens            = 6
            output_cost           = 6.56
            temperature           = 3.27
            top_k                 = 2
            top_p                 = 0.5
            upstream_url          = "https://baggy-trash.biz/"
            deployment_id         = "claude-5-model"
          }
        }
        name     = "claude-5-model"
        provider = konnect_ai_gateway_model_provider.my_aigatewaymodelprovider_anthropic.name
      }
    ]
  }
  gateway_id = konnect_ai_gateway.my_aigateway.id
}
