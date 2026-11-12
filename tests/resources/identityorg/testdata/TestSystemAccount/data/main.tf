data "konnect_system_account" "my_systemaccount" {
  filter = {
    name = {
      eq = "Sample System Account"
    }
  }
}

data "konnect_system_account_list" "my_systemaccountlist" {
  filter = {
    name = {
      eq = "Sample System Account"
    }
  }
}

output "system_account_id" {
  value = data.konnect_system_account.my_systemaccount.id
}

output "system_accounts" {
  value = data.konnect_system_account_list.my_systemaccountlist
}
