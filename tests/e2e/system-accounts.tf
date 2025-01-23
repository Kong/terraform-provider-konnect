resource "konnect_system_account" "my_systemaccount" {
  name            = "Michael Sat Test"
  description     = "Demo System Account"
  konnect_managed = false
}

resource "konnect_system_account_access_token" "my_systemaccountaccesstoken" {
  name       = "TF Managed Token"
  expires_at = "2025-12-31T23:59:59Z"
  account_id = konnect_system_account.my_systemaccount.id
}

resource "konnect_system_account_role" "my_systemaccountrole" {
  entity_id        = "*"
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Viewer"
  account_id       = konnect_system_account.my_systemaccount.id
}

resource "konnect_system_account_team" "my_systemaccountteam" {
  account_id = konnect_system_account.my_systemaccount.id
  team_id    = konnect_team.my_team.id
}
