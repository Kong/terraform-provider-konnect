# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)
#
# A role assignment can be scoped to a whole control plane (a one element
# entity_ids) or to a single workspace inside it (entity_ids with two ids).
# Workspace roles such as "Workspace Admin", "Workspace Viewer" and
# "Workspace Deployer" take two positional ids, parent first:
#
#   entity_ids = [<control_plane_id>, <workspace_id>]
#
# entity_ids is ordered, so the provider sends the list as is. entity_id and
# entity_ids are mutually exclusive; the provider rejects a configuration that
# sets both.

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

# Read access to every workspace in the control plane, for contrast. One id
# means the control plane itself.
#
# Keep the attribute set when you want to change the scope: removing entity_ids
# from an existing assignment leaves the assignment as it is, because a null
# configuration value is not a diff there. Replacing entity_ids with the
# deprecated entity_id does the same: the assignment keeps its old scope.
resource "konnect_team_role" "payments_control_plane_viewer" {
  team_id          = konnect_team.payments.id
  role_name        = "Viewer"
  entity_type_name = "Control Planes"
  entity_region    = "us"

  entity_ids = [konnect_gateway_control_plane.tfdemo.id]
}
