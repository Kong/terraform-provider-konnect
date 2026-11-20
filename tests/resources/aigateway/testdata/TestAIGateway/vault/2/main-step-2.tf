resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - vault"
  name         = "tf-test-aigw-vault"
}

resource "konnect_ai_gateway_config_store" "my_aigatewayconfigstore" {
  force      = true
  gateway_id = konnect_ai_gateway.my_aigateway.id
  name       = "tf-test-config-store"
}

resource "konnect_ai_gateway_vault" "my_aigatewayvault" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  konnect = {
    config = {
      config_store_id = konnect_ai_gateway_config_store.my_aigatewayconfigstore.id
    }
    name = "tf-test-vault-updated"
  }
}
