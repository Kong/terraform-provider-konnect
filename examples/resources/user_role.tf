resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "My CP"

  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_team" "my_team" {
  description = "My team description."

  name = "MyTeamName"
}

resource "konnect_team_user" "my_teamuser" {
  team_id = konnect_team.my_team.id
  user_id = "my-user-id"
}

resource "konnect_user_role" "my_userrole" {
  entity_id        = konnect_gateway_control_plane.my_konnect_cp.id
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Viewer"
  user_id          = konnect_team_user.my_teamuser.user_id
}
