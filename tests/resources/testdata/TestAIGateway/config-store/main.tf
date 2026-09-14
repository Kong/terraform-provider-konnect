resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - config-store"
  name         = "tf-test-aigw-config-store"
}

resource "konnect_ai_gateway_config_store" "my_aigatewayconfigstore" {
  provider   = konnect-beta
  force      = true
  gateway_id = konnect_ai_gateway.my_aigateway.id
  name       = "tf-test-config-store"
}
