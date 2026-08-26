data "konnect_team_list" "my_teamlist" {
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

output "team_list" {
  value = data.konnect_team_list.my_teamlist
}
