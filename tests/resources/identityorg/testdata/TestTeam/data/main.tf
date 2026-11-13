data "konnect_team" "my_team" {
  filter = {
    name = {
      eq = "portal-admin"
    }
  }
}

data "konnect_team_list" "my_teamlist" {
  filter = {
    name = {
      eq = "portal-admin"
    }
  }
}

output "team_id" {
  value = data.konnect_team.my_team.id
}

output "teams" {
  value = data.konnect_team_list.my_teamlist
}
