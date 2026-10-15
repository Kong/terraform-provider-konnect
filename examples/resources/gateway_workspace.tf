resource "konnect_gateway_workspace" "my_gatewayworkspace" {
  name             = "team-1"
  description      = "A test workspace for team 1"
  comment          = "A test workspace for team 1"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
