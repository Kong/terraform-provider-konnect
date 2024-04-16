resource "konnect_gateway_upstream" "my_upstream" {
  name             = "demo-upstream"
  slots            = 30
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
