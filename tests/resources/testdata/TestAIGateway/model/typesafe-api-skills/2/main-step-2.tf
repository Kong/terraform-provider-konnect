resource "konnect_ai_gateway" "my_aigateway" {
  display_name        = "TF Test AIGW - model-typesafe-api-skills"
  name                = "tf-test-aigw-model-typesafe-api-skills"
  min_runtime_version = "2.2"
}

resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider_typesafe_api_skills" {
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
    display_name = "TF Test Typesafe AI Provider - api skills"
    name         = "tf-test-typesafe-api-skills-provider"
  }
}

resource "konnect_ai_gateway_model" "my_aigatewaymodel_typesafe_api_skills" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  api = {
    capabilities = [
      "skills"
    ]
    config = {
      route = {
        hosts = []
        paths = [
          "/my-typesafe-api-skills-test-path"
        ]
        https_redirect_status_code = 426
        preserve_host              = false
        regex_priority             = 0
        request_buffering          = true
        response_buffering         = true
        strip_path                 = true
      }
    }
    display_name = "Skills Capability Test - Typesafe (api) Updated"
    enabled      = true
    formats = [
      {
        type = "passthrough"
      }
    ]
    name = "tf-test-typesafe-api-skills-model"
    targets = [
      {
        config = {
          typesafe = {
            input_cost   = 3.7
            output_cost  = 6.56
            upstream_url = "https://baggy-trash.biz/"
          }
        }
        name     = "typesafe-api-skills-model"
        provider = konnect_ai_gateway_model_provider.my_aigatewaymodelprovider_typesafe_api_skills.name
      }
    ]
  }
}
