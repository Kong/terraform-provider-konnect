resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - consumer-credential"
  name         = "tf-test-aigw-consumer-credential"
}

resource "konnect_ai_gateway_consumer" "my_aigatewayconsumer" {
  display_name = "TF Test Consumer"
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name         = "tf-test-consumer"
  type         = "api-key"
}

resource "konnect_ai_gateway_consumer_credential" "my_aigatewayconsumercredential" {
  api_key      = "tf-test-dev-key"
  consumer_id  = konnect_ai_gateway_consumer.my_aigatewayconsumer.id
  display_name = "TF Test Dev Key"
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name         = "tf-test-dev-key"
  type         = "api-key"
}
