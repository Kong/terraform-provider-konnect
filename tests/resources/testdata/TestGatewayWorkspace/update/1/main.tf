resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Gateway Workspace Update"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "my_gatewayworkspace" {
  name        = "tf-acceptance-workspace-update"
  description = "A test workspace for team 1"
  comment     = "A test workspace for team 1"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
