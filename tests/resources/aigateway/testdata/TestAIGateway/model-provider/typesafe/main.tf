resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - model-provider-typesafe"
  name         = "tf-test-aigw-model-provider-typesafe"
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
