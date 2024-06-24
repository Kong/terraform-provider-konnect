# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)

resource "konnect_team" "my_team" {
  name        = "My Terraform Team"
  description = "This is a team that is managed by Terraform"

  labels = {
    example = "here"
  }
}

resource "konnect_team_role" "my_team_role" {
  entity_id        = "*"
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Admin"
  team_id          = konnect_team.my_team.id
}

resource "konnect_team_user" "my_team_user" {
  # The UUID of the user to add to the team
  # Make sure to change this to the UUID of the user you want to add
  user_id = "712675bd-asap-4263-892b-90dbc5a22576"
  team_id = konnect_team.my_team.id
}
