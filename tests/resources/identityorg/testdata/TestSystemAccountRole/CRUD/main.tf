resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For System account role"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_system_account" "my_systemaccount" {
  description     = "my description"
  konnect_managed = false
  name            = "TFAcceptanceSystemAccountName"
}


resource "konnect_system_account_role" "my_systemaccountrole" {
  entity_id        = konnect_gateway_control_plane.my_konnect_cp.id
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Viewer"
  account_id       = konnect_system_account.my_systemaccount.id
}
