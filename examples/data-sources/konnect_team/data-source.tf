data "konnect_team" "my_team" {
  filter = {
    name = {
      contains = "...my_contains..."
      eq       = "...my_eq..."
    }
  }
}