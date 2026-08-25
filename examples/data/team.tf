data "konnect_team" "my_team" {
  filter = {
    name = {
      eq = "portal-admin"
    }
    labels = {
      example = {
        eq = "here"
      }
    }
  }
}

output "team_id" {
  value = data.konnect_team.my_team.id
}
