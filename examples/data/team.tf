data "konnect_team" "my_team" {
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

output "team_id" {
  value = data.konnect_team.my_team.id
}
