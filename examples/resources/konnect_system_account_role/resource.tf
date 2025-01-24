resource "konnect_system_account_role" "my_systemaccountrole" {
  account_id       = "...my_account_id..."
  entity_id        = "e67490ce-44dc-4cbd-b65e-b52c746fc26a"
  entity_region    = "eu"
  entity_type_name = "Control Planes"
  role_name        = "Viewer"
}