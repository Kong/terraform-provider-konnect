data "konnect_team" "my_team" {
  filter = {
    name = {
      eq = "portal-admin"
    }
    labels = {
      my_label_key = {
        eq = "here"
      }
    }
  }
}

output "team_id" {
  value = data.konnect_team.my_team.id
}
