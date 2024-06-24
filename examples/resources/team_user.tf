resource "konnect_team_user" "my_team_user" {
  # The UUID of the user to add to the team
  # Make sure to change this to the UUID of the user you want to add
  user_id = "712675bd-asap-4263-892b-90dbc5a22576"
  team_id = konnect_team.my_team.id
}
