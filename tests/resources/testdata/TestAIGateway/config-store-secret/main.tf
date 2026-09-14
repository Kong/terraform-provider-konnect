resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - config-store-secret"
  name         = "tf-test-aigw-config-store-secret"
}

resource "konnect_ai_gateway_config_store" "my_aigatewayconfigstore" {
  provider   = konnect-beta
  force      = true
  gateway_id = konnect_ai_gateway.my_aigateway.id
  name       = "tf-test-config-store"
}

resource "konnect_ai_gateway_config_store_secret" "my_aigatewayconfigstoresecret" {
  provider        = konnect-beta
  config_store_id = konnect_ai_gateway_config_store.my_aigatewayconfigstore.id
  gateway_id      = konnect_ai_gateway.my_aigateway.id
  key             = "tf-test-secret-key"
  value           = "tf-test-secret-value"
}
