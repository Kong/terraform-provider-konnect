data "konnect_team" "my_team" {
  filter = {
    name = {
      eq = "portal-admin"
    }
  }
}

output "team_id" {
  value = data.konnect_team.my_team.id
}
