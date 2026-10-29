resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - datastore token-vault"
  name         = "tf-test-aigw-datastore-token-vault"
}

resource "konnect_ai_gateway_datastore" "my_aigatewaydatastore_token_vault" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  token_vault = {
    name         = "tf-test-token-vault"
    display_name = "TF Test Token Vault"
    description  = "Token vault for TF acceptance tests"
    config = {
      directory       = "tf-test-tokens"
      current_version = "v2"
      keys = [
        {
          key     = "{vault://tf-test-vault/token-vault-key-v1}"
          version = "v1"
        },
        {
          key     = "{vault://tf-test-vault/token-vault-key-v2}"
          version = "v2"
        }
      ]
    }
  }
}
