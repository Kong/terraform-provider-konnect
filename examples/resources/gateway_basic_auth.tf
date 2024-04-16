resource "konnect_gateway_basic_auth" "my_basicauth" {
  username = "alice"
  password = "demo"

  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
