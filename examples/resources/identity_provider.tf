# OIDC
resource "konnect_identity_provider" "my_identityprovider" {
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
# SAML
resource "konnect_identity_provider" "my_identityprovider" {
  enabled    = true
  type       = "saml"
  login_path = "testsaml"
  config = {
    saml_identity_provider_config = {
      idp_metadata_url = "https://mocksaml.com/api/saml/metadata"
      idp_metadata_xml = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"
    }
  }
}
