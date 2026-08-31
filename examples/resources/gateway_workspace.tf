resource "konnect_gateway_workspace" "my_gatewayworkspace" {
  name        = "team-1"
  description = "A test workspace for team 1"
  comment     = "A test workspace for team 1"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Gateway entities are created in the "default" workspace unless the optional
# `workspace` argument names another one. Changing it recreates the entity.
resource "konnect_gateway_service" "team_1_httpbin" {
  name     = "HTTPBin"
  protocol = "https"
  host     = "httpbin.org"
  port     = 443
  path     = "/"

  workspace        = konnect_gateway_workspace.my_gatewayworkspace.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
