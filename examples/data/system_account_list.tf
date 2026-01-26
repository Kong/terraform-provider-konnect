data "konnect_system_account_list" "my_systemaccountlist" {
  filter = {
    name = {
      eq = "Sample System Account"
    }
  }
}

output "system_accounts" {
  value = data.konnect_system_account_list.my_systemaccountlist
}
