resource "konnect_team_role" "my_team_role" {
  entity_id        = "*"
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Admin"
  team_id          = konnect_team.my_team.id
}