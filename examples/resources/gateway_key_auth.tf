resource "konnect_gateway_key_auth" "my_keyauth" {
  key              = "abc123"
  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
