resource "konnect_gateway_acl" "my_acl" {
  group = "internal_users"

  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
