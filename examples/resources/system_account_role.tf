resource "konnect_system_account_role" "my_systemaccountrole" {
  entity_id        = "*"
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Viewer"
  account_id       = konnect_system_account.my_systemaccount.id
}
