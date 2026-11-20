resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - datastore redis-ce"
  name         = "tf-test-aigw-datastore-redis-ce"
}

resource "konnect_ai_gateway_datastore" "my_aigatewaydatastore_redis_ce" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  redis_ce = {
    name         = "tf-test-redis-ce"
    display_name = "TF Test Redis CE"
    config = {
      host = "127.0.0.1"
      port = 6379
    }
  }
}
