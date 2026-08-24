data "konnect_team_list" "my_teamlist" {
  filter = {
    name = {
      eq = "portal-admin"
    }
    labels = {
      example = {
        legacy_string_field_filter = {
          eq = "here"
        }
      }
    }
  }
}

output "team_list" {
  value = data.konnect_team_list.my_teamlist
}
