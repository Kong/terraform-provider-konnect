resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - datastore redis-ee"
  name         = "tf-test-aigw-datastore-redis-ee"
}

resource "konnect_ai_gateway_datastore" "my_aigatewaydatastore_redis_ee" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  redis_ee = {
    name         = "tf-test-redis-ee"
    display_name = "TF Test Redis EE"
    config = {
      host = "127.0.0.1"
      port = {
        integer = 6379
      }
    }
  }
}
