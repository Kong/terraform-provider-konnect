# SAML
resource "konnect_identity_provider" "saml_provider" {
  enabled    = true
  type       = "saml"
  login_path = "testsaml"
  config = {
    saml_identity_provider_config = {
      idp_metadata_url = "https://mocksaml.com/api/saml/metadata"
    }
  }
}
