resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Workspace Move"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "target" {
  name        = "tf-acceptance-workspace-move-target"
  description = "Destination workspace for the move test"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Step 1: no `workspace` in config, so this lands in "default".
resource "konnect_gateway_service" "httpbin" {
  name     = "HTTPBin"
  protocol = "https"
  host     = "httpbin.org"
  port     = 443
  path     = "/"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
