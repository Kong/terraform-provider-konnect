# An OIDC issuer this directory trusts. Principals are provisioned on first
# exchange, using the claims the mapping names.
resource "konnect_identity_trusted_idp" "my_trusted_idp" {
  directory_id = konnect_identity_directory.my_directory.id
  display_name = "Corporate Okta"
  issuer_url   = "https://acme.okta.com/oauth2/default"

  jit_provisioning_enabled = true

  # Allow principals provisioned by this IdP to exchange their token for a
  # token vault credential.
  vault_access_enabled = true

  jit = {
    mapping = {
      display_name = "name"
      email        = "email"
      groups       = "groups"
    }
  }
}
