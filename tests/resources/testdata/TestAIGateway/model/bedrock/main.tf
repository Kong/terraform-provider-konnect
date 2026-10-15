resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - model-bedrock"
  name         = "tf-test-aigw-model-bedrock"
}

resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider_bedrock" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  bedrock = {
    config = {
      auth = {
        aws = {}
      }
    }
    display_name = "TF Test Bedrock AI Provider"
    name         = "tf-test-bedrock-provider"
  }
}

resource "konnect_ai_gateway_model" "my_aigatewaymodel_bedrock" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  model = {
    capabilities = [
      "generate"
    ]
    config = {
      route = {
        paths = [
          "/my-bedrock-base-path"
        ]
      }
    }
    display_name = "My Test Bedrock model"
    enabled      = true
    formats = [
      {
        type = "openai"
      }
    ]
    name = "my-bedrock-model"
    targets = [
      {
        # region/max_tokens only: batch_bucket_prefix, performance_config_latency
        # and video_output_s3_uri are intentionally left unset here, regression-testing
        # the AIGW overlay fix for BedrockModelConfig (these previously had to be sent
        # as explicit null and were rejected by the API as "not nullable").
        config = {
          bedrock = {
            max_tokens = 1024
            region     = "us-west-2"
          }
        }
        name     = "us.anthropic.claude-haiku-4-5-20251001-v1:0"
        provider = konnect_ai_gateway_model_provider.my_aigatewaymodelprovider_bedrock.name
      }
    ]
  }
}
