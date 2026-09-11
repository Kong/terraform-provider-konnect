resource "konnect_identity_provider" "saml_provider_tf" {
  enabled    = false
  type       = "saml"
  login_path = "testsamldatatf"
  config = {
    saml_identity_provider_config = {
      idp_metadata_url = "https://mocksaml.com/api/saml/metadata"
    }
  }
}

data "konnect_identity_provider" "by_id" {
  id = konnect_identity_provider.saml_provider_tf.id
}

output "idp_login_path" {
  value = data.konnect_identity_provider.by_id.login_path
}