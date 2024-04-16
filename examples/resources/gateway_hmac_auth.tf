resource "konnect_gateway_hmac_auth" "my_hmac" {
  username         = "alice"
  secret           = "secret1234"
  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
