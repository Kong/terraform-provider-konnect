resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - datastore token-vault literal key"
  name         = "tf-test-aigw-datastore-token-vault-literal-key"
}

resource "konnect_ai_gateway_datastore" "my_aigatewaydatastore_token_vault_literal_key" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  token_vault = {
    name         = "tf-test-token-vault-literal-key"
    display_name = "TF Test Token Vault Literal Key"
    config = {
      directory       = "tf-test-tokens"
      current_version = "v1"
      keys = [
        {
          key     = "tftestliteralkeyvaluenotavaultreference"
          version = "v1"
        }
      ]
    }
  }
}
