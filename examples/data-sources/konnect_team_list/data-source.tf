data "konnect_team_list" "my_teamlist" {
  filter = {
    name = {
      contains = "...my_contains..."
      eq       = "...my_eq..."
    }
  }
}