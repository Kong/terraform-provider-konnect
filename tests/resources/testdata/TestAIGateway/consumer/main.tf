resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - consumer"
  name         = "tf-test-aigw-consumer"
}

resource "konnect_ai_gateway_consumer" "my_aigatewayconsumer" {
  provider     = konnect-beta
  display_name = "TF Test Consumer"
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name         = "tf-test-consumer"
  type         = "api-key"
}
