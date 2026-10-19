resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Gateway Workspace DS"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "ds" {
  name        = "tf-acceptance-workspace-ds"
  description = "A test workspace for the data source"
  comment     = "A test workspace for the data source"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

data "konnect_gateway_workspace" "by_name" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  filter_name_eq   = konnect_gateway_workspace.ds.name

  depends_on = [konnect_gateway_workspace.ds]
}

output "workspace_id" {
  value = data.konnect_gateway_workspace.by_name.id
}

output "workspace_name" {
  value = data.konnect_gateway_workspace.by_name.name
}
