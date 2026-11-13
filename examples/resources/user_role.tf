resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name = "My CP"

  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "my_workspace" {
  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
  name             = "my-workspace"
}

resource "konnect_team" "my_team" {
  description = "My team description."

  name = "MyTeamName"
}

resource "konnect_team_user" "my_teamuser" {
  team_id = konnect_team.my_team.id
  user_id = "my-user-id"
}

# A whole control plane: one id.
resource "konnect_user_role" "my_userrole" {
  entity_ids       = [konnect_gateway_control_plane.my_konnect_cp.id]
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Viewer"
  user_id          = konnect_team_user.my_teamuser.user_id
}

# A role that acts inside a single workspace takes two ids, parent first:
# [<control_plane_id>, <workspace_id>]
resource "konnect_user_role" "my_userrole_workspace" {
  entity_ids = [
    konnect_gateway_control_plane.my_konnect_cp.id,
    konnect_gateway_workspace.my_workspace.id,
  ]
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Workspace Viewer"
  user_id          = konnect_team_user.my_teamuser.user_id
}
