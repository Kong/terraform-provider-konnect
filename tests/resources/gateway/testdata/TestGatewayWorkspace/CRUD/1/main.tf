resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Gateway Workspace"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "my_gatewayworkspace" {
  name        = "tf-acceptance-workspace"
  description = "A test workspace for team 1"
  comment     = "A test workspace for team 1"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Created in the workspace above rather than the control plane's "default"
# workspace, so the workspace-scoped entity paths get exercised too.
resource "konnect_gateway_service" "httpbin" {
  name     = "HTTPBin"
  protocol = "https"
  host     = "httpbin.org"
  port     = 443
  path     = "/"

  workspace        = konnect_gateway_workspace.my_gatewayworkspace.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
