resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - config-store"
  name         = "tf-test-aigw-config-store"
}

resource "konnect_ai_gateway_config_store" "my_aigatewayconfigstore" {
  force      = false
  gateway_id = konnect_ai_gateway.my_aigateway.id
  name       = "tf-test-config-store"
}
