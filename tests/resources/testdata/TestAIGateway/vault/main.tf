resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - vault"
  name         = "tf-test-aigw-vault"
}

resource "konnect_ai_gateway_config_store" "my_aigatewayconfigstore" {
  provider   = konnect-beta
  force      = true
  gateway_id = konnect_ai_gateway.my_aigateway.id
  name       = "tf-test-config-store"
}

resource "konnect_ai_gateway_vault" "my_aigatewayvault" {
  provider   = konnect-beta
  gateway_id = konnect_ai_gateway.my_aigateway.id
  konnect = {
    config = {
      config_store_id = konnect_ai_gateway_config_store.my_aigatewayconfigstore.id
    }
    name = "tf-test-vault"
  }
}
