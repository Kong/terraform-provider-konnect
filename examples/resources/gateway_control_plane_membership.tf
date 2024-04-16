resource "konnect_gateway_control_plane_membership" "my_gatewaycontrolplanemembership" {
  id = konnect_gateway_control_plane.tfcpg.id
  members = [
    { id = konnect_gateway_control_plane.tfdemo.id },
    { id = konnect_gateway_control_plane.tfmember.id },
  ]
}
