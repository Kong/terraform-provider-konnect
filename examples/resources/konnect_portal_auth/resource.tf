resource "konnect_portal_auth" "my_portalauth" {
  basic_auth_enabled      = true
  idp_mapping_enabled     = true
  konnect_mapping_enabled = false
  oidc_auth_enabled       = false
  oidc_claim_mappings = {
    email  = "email"
    groups = "custom-group-claim"
    name   = "name"
  }
  oidc_client_id     = "x7id0o42lklas0blidl2"
  oidc_client_secret = "...my_oidc_client_secret..."
  oidc_issuer        = "https://identity.example.com/v2"
  oidc_scopes = [
    "email",
    "openid",
    "profile",
  ]
  oidc_team_mapping_enabled = true
  portal_id                 = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  saml_auth_enabled         = false
}