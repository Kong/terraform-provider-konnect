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

data "konnect_identity_provider" "my_identity_provider" {
  id = konnect_identity_provider.saml_provider_tf.id
}

data "konnect_identity_provider_list" "my_list_identity_provider" {
  filter = {
    type = {
        eq = "saml"
      }
    }
  depends_on = [konnect_identity_provider.saml_provider_tf]
}

output "identity_provider_list" {
  value = data.konnect_identity_provider_list.my_list_identity_provider
}

output "identity_provider" {
  value = data.konnect_identity_provider.my_identity_provider
}