resource "konnect_dcr_provider" "my_dcrprovider" {
  kong_identity = {
    dcr_config    = {}
    issuer        = "https://issuer.example.com"
    name          = "my-dcr-provider"
    provider_type = "kongIdentity"
    labels = {
      team = "platform"
    }
  }
}

resource "konnect_application_auth_strategy" "my_authstrategy" {
  openid_connect = {
    strategy_type = "openid_connect"
    name          = "my-auth-strategy"
    display_name  = "My Auth Strategy"

    dcr_provider_id = konnect_dcr_provider.my_dcrprovider.id

    configs = {
      openid_connect = {
        issuer           = "https://issuer.example.com"
        credential_claim = ["sub"]
        auth_methods     = ["client_credentials"]
        scopes           = ["openid"]
      }
    }
  }
}
