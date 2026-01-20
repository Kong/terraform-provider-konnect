data "konnect_team_list" "my_teamlist" {
  filter = {
    name = {
      eq = "portal-admin"
    }
  }
}

output "team_list" {
  value = data.konnect_team_list.my_teamlist
}
