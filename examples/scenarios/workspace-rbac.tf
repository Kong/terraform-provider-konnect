# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)
#
# Workspace scoped roles take two positional ids, parent first:
# entity_ids = [<control_plane_id>, <workspace_id>]

resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Workspace RBAC demo"
  description  = "Control plane for the workspace RBAC scenario"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "payments" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  name             = "payments"
  description      = "Workspace for the payments team"
}

resource "konnect_team" "payments" {
  name        = "Payments"
  description = "Team that owns the payments workspace"
}

# Admin inside the payments workspace only.
resource "konnect_team_role" "payments_workspace_admin" {
  team_id          = konnect_team.payments.id
  role_name        = "Workspace Admin"
  entity_type_name = "Control Planes"
  entity_region    = "us"

  entity_ids = [
    konnect_gateway_control_plane.tfdemo.id,
    konnect_gateway_workspace.payments.id,
  ]
}

# Control plane wide read access, for contrast.
# Keep entity_ids set to change the scope: a null value is not a diff.
resource "konnect_team_role" "payments_control_plane_viewer" {
  team_id          = konnect_team.payments.id
  role_name        = "Viewer"
  entity_type_name = "Control Planes"
  entity_region    = "us"

  entity_ids = [konnect_gateway_control_plane.tfdemo.id]
}
