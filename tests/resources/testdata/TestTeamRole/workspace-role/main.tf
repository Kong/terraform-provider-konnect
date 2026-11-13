resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For team role workspace"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "team_payments" {
  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
  name             = "team-payments"
  description      = "Workspace for the payments team"
}

resource "konnect_team" "my_team" {
  description = "TF acceptance test team description."

  name = "TFAcceptanceTeamRoleWorkspace"
}

# Workspace scoped roles take two positional ids: the control plane, then the
# workspace. "Workspace Admin" is not in the role_name enum, so it is passed
# through because role_name allows unknown values.
resource "konnect_team_role" "my_teamrole" {
  entity_ids = [
    konnect_gateway_control_plane.my_konnect_cp.id,
    konnect_gateway_workspace.team_payments.id,
  ]
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Workspace Admin"
  team_id          = konnect_team.my_team.id
}
