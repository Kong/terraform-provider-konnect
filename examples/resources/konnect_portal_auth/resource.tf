resource "konnect_portal_auth" "my_portalauth" {
  basic_auth_enabled      = true
  konnect_mapping_enabled = false
  oidc_auth_enabled       = false
  oidc_claim_mappings = {
    email  = "email"
    groups = "custom-group-claim"
    name   = "name"
  }
  oidc_client_id     = "...my_oidc_client_id..."
  oidc_client_secret = "...my_oidc_client_secret..."
  oidc_issuer        = "...my_oidc_issuer..."
  oidc_scopes = [
    "...",
  ]
  oidc_team_mapping_enabled = true
  portal_id                 = "ee74cf0f-d04e-4faf-8ab4-9554032d0057"
}