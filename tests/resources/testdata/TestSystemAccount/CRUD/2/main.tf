resource "konnect_system_account" "my_systemaccount" {
  description     = "my updated description"
  konnect_managed = false
  name            = "TFAcceptanceSystemAccountName"
}