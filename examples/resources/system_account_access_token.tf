resource "konnect_system_account_access_token" "my_systemaccountaccesstoken" {
  name       = "TF Managed Token"
  expires_at = "2024-12-31T23:59:59Z"
  account_id = konnect_system_account.my_systemaccount.id
}
