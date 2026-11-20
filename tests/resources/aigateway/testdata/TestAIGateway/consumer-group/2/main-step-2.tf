resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - consumer-group"
  name         = "tf-test-aigw-consumer-group"
}

resource "konnect_ai_gateway_consumer_group" "my_aigatewayconsumergroup" {
  display_name = "TF Test Consumer Group Updated"
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name         = "tf-test-consumers"
}
