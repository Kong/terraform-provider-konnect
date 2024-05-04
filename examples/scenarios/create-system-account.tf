# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)

resource "konnect_system_account" "my_systemaccount" {
  name            = "My System Account"
  description     = "Demo System Account"
  konnect_managed = false
}

resource "konnect_system_account_access_token" "my_systemaccountaccesstoken" {
  name       = "TF Managed Token"
  expires_at = "2024-12-31T23:59:59Z"
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

  # Update this to your team ID manually until the Teams data source is available
  # You can find this using the API
  # GET https://global.api.konghq.com/v3/teams
  team_id = "3e3a7539-4d61-4424-a714-fdfdc3728fad"
}

output "token" {
  value = konnect_system_account_access_token.my_systemaccountaccesstoken.token
}
