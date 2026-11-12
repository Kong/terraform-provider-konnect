resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For user role"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_team" "my_team" {
  description = "TF acceptance test team description."

  name = "TFAcceptanceUserRoleTeamName"
}

resource "konnect_team_user" "my_teamuser" {
  team_id = konnect_team.my_team.id
  user_id = "df5404e1-5c36-414b-a40b-903387226475" // creating user via TF is not possible due to email verification requirement.
}

resource "konnect_user_role" "my_userrole" {
  entity_id        = konnect_gateway_control_plane.my_konnect_cp.id
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Viewer"
  user_id          = konnect_team_user.my_teamuser.user_id
}
