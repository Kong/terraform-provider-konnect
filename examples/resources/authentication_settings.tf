resource "konnect_authentication_settings" "my_authenticationsettings" {
  # Login options
  oidc_auth_enabled       = false
  saml_auth_enabled       = false
  basic_auth_enabled      = true

  # Team mapping
  idp_mapping_enabled     = true
  konnect_mapping_enabled = false
}
