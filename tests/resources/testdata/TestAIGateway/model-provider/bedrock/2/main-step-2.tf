resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - model-provider-bedrock"
  name         = "tf-test-aigw-model-provider-bedrock"
}

resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider_bedrock" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  bedrock = {
    config = {
      auth = {
        aws = {}
      }
    }
    display_name = "TF Test Bedrock AI Provider Updated"
    name         = "tf-test-bedrock-provider"
  }
}
