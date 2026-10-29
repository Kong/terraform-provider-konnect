resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Workspace ManagedBy"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

# `managed_by` is Optional + Computed, which is the shape that produces a
# perpetual diff if the API echoes back anything other than what was sent
# (extra keys, reordering, normalised values). Setting it in config and then
# re-planning the same config is what catches that.
resource "konnect_gateway_workspace" "my_gatewayworkspace" {
  name        = "tf-acceptance-workspace-managed"
  description = "Workspace with managed_by set"

  managed_by = {
    terraform = "true"
    team      = "platform"
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
