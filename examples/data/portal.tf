data "konnect_portal" "my_portal" {
  filter = {
    name = {
      eq = "My Portal"
    }
  }
}
