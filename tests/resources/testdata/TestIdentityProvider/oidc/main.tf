# OIDC
resource "konnect_identity_provider" "oidc_provider" {
  enabled    = "true"
  type       = "oidc"
  login_path = "testoidc"
  config = {
    oidc_identity_provider_config = {
      client_id     = "client-id"
      client_secret = "my-client-secret"
      issuer_url    = "https://accounts.google.com"
      scopes = [
        "openid",
        "email",
        "profile"
      ]
      claim_mappings = {
        email  = "email"
        groups = "groups"
        name   = "name"
      }
    }
  }
}
