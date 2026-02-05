data "konnect_system_account" "my_systemaccount" {
  filter = {
    description = {
      contains = "...my_contains..."
      eq       = "...my_eq..."
    }
    konnect_managed = true
    name = {
      contains = "...my_contains..."
      eq       = "...my_eq..."
    }
  }
}