resource "konnect_system_account" "my_systemaccount" {
  description     = "my description"
  konnect_managed = false
  name            = "TFAcceptanceSystemAccountName"
}