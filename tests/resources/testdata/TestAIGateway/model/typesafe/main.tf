resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - model-typesafe"
  name         = "tf-test-aigw-model-typesafe"
  min_runtime_version = "2.2"
}

resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider_typesafe" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  typesafe = {
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
    display_name = "TF Test Typesafe AI Provider"
    name         = "tf-test-typesafe-provider"
  }
}

resource "konnect_ai_gateway_model" "my_aigatewaymodel_model" {
  model = {
    capabilities = [
      "generate"
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
    display_name = "My Test Typesafe model"
    enabled      = true
    formats = [
      {
        type = "openai"
      }
    ]
    name = "tf-test-typesafe-model"
    targets = [
      {
        config = {
          typesafe = {
            input_cost   = 3.7
            output_cost  = 6.56
            upstream_url = "https://baggy-trash.biz/"
          }
        }
        name     = "typesafe-model"
        provider = konnect_ai_gateway_model_provider.my_aigatewaymodelprovider_typesafe.name
      }
    ]
  }
  gateway_id = konnect_ai_gateway.my_aigateway.id
}
