resource "konnect_gateway_consumer_group" "gold" {
  name             = "gold"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
