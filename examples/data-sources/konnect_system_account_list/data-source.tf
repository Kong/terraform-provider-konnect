data "konnect_system_account_list" "my_systemaccountlist" {
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