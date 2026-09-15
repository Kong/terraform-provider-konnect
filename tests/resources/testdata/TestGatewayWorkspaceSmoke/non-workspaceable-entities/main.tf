resource "konnect_gateway_control_plane" "smoke_nonworkspaceable_cp" {
  name         = "Terraform Control Plane For Non-Workspaceable Smoke Test"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

# Present to prove non-workspaceable entities are unaffected by, and don't
# require, a workspace even when one exists on the same control plane.
resource "konnect_gateway_workspace" "smoke" {
  name        = "tf-acceptance-smoke-workspace"
  description = "A workspace that coexists with non-workspaceable entities"

  control_plane_id = konnect_gateway_control_plane.smoke_nonworkspaceable_cp.id
}

# A representative sample of commonly used entities that are NOT
# workspaceable: they are always addressed directly through the control
# plane and have no "workspace" attribute at all.

resource "konnect_gateway_config_store" "smoke" {
  name = "smoke-config-store"

  control_plane_id = konnect_gateway_control_plane.smoke_nonworkspaceable_cp.id
}

resource "konnect_gateway_config_store_secret" "smoke" {
  key   = "smoke-key"
  value = "smoke-value"

  config_store_id  = konnect_gateway_config_store.smoke.id
  control_plane_id = konnect_gateway_control_plane.smoke_nonworkspaceable_cp.id
}
