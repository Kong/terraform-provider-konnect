resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For team role entity ids"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_team" "my_team" {
  description = "TF acceptance test team description."

  name = "TFAcceptanceTeamRoleEntityIds"
}

resource "konnect_team_role" "my_teamrole" {
  entity_ids       = [konnect_gateway_control_plane.my_konnect_cp.id]
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Viewer"
  team_id          = konnect_team.my_team.id
}
