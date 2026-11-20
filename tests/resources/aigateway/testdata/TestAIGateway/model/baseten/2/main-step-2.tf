resource "konnect_ai_gateway" "my_aigateway" {
  display_name        = "TF Test AIGW - model-baseten"
  name                = "tf-test-aigw-model-baseten"
  min_runtime_version = "2.3"
}

resource "konnect_ai_gateway_datastore" "my_aigatewaydatastore_token_vault" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  token_vault = {
    name         = "tf-test-baseten-token-vault"
    display_name = "TF Test Baseten Token Vault"
    config = {
      directory       = "tf-test-baseten-tokens"
      current_version = "v1"
      keys = [
        {
          key     = "{vault://tf-test-vault/baseten-key-v1}"
          version = "v1"
        }
      ]
    }
  }
}

resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider_baseten" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  baseten = {
    config = {
      auth = {
        headers = []
        params  = []
      }
    }
    datastores = [
      {
        name = konnect_ai_gateway_datastore.my_aigatewaydatastore_token_vault.token_vault.name
      }
    ]
    display_name = "TF Test Baseten AI Provider Updated"
    name         = "tf-test-baseten-provider"
    token_vault = {
      provider = "tf-test-baseten"
    }
  }
}

resource "konnect_ai_gateway_model" "my_aigatewaymodel_baseten" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  model = {
    capabilities = [
      "generate"
    ]
    config = {
      route = {
        paths = [
          "/my-baseten-base-path"
        ]
      }
    }
    display_name = "My Test Baseten model Updated"
    enabled      = true
    formats = [
      {
        type = "openai"
      }
    ]
    name = "my-baseten-model"
    targets = [
      {
        config = {
          baseten = {
            max_tokens = 1024
          }
        }
        name     = "deepseek-ai/DeepSeek-V4-Pro-0813"
        provider = konnect_ai_gateway_model_provider.my_aigatewaymodelprovider_baseten.name
      }
    ]
  }
}
