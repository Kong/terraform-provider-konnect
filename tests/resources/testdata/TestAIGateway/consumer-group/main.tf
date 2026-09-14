resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - consumer-group"
  name         = "tf-test-aigw-consumer-group"
}

resource "konnect_ai_gateway_consumer_group" "my_aigatewayconsumergroup" {
  provider     = konnect-beta
  display_name = "TF Test Consumer Group"
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name         = "tf-test-consumers"
}
