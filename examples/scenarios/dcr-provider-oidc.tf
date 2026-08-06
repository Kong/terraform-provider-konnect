# Dynamic Client Registration (DCR) with an OpenID Connect application auth
# strategy.
#
# konnect_dcr_provider lets Konnect auto-register OAuth/OIDC applications
# with an identity provider on behalf of developers, instead of requiring
# them to be manually pre-registered with the IdP. Linking a DCR provider to
# a konnect_application_auth_strategy (via dcr_provider_id) means any
# application registered against that auth strategy gets its IdP client
# created automatically through DCR.
#
# konnect_dcr_provider is a "one-of" resource -- exactly one variant block
# (kong_identity, auth0, azure_ad, curity, okta, http) is set at a time.
# This example uses kong_identity, which needs no external identity
# provider: Konnect itself acts as the DCR authority. The commented blocks
# below show how to configure every other supported variant instead --
# uncomment exactly one (and remove kong_identity) to switch.

resource "konnect_dcr_provider" "my_dcrprovider" {
  kong_identity = {
    name         = "my-dcr-provider"
    display_name = "My DCR Provider"
    issuer       = "https://issuer.example.com"

    dcr_config = {}

    labels = {
      team = "platform"
    }
  }

  # --- Auth0 variant (needs a real Auth0 management client) ---
  # auth0 = {
  #   name          = "my-dcr-provider-auth0"
  #   display_name  = "My DCR Provider (Auth0)"
  #   issuer        = "https://my-tenant.us.auth0.com"
  #
  #   dcr_config = {
  #     initial_client_id            = var.auth0_initial_client_id
  #     initial_client_secret        = var.auth0_initial_client_secret
  #     initial_client_audience      = "https://my-tenant.us.auth0.com/api/v2/"
  #     use_developer_managed_scopes = false
  #   }
  # }

  # --- Azure AD variant (needs an app registration in the Azure tenant) ---
  # azure_ad = {
  #   name          = "my-dcr-provider-azuread"
  #   display_name  = "My DCR Provider (Azure AD)"
  #   issuer        = "https://login.microsoftonline.com/${var.azure_ad_tenant_id}/v2.0"
  #
  #   dcr_config = {
  #     initial_client_id     = var.azure_ad_initial_client_id
  #     initial_client_secret = var.azure_ad_initial_client_secret
  #   }
  # }

  # --- Curity variant (needs a management client in the Curity server) ---
  # curity = {
  #   name          = "my-dcr-provider-curity"
  #   display_name  = "My DCR Provider (Curity)"
  #   issuer        = "https://curity.example.com/oauth/v2/oauth-anonymous"
  #
  #   dcr_config = {
  #     initial_client_id     = var.curity_initial_client_id
  #     initial_client_secret = var.curity_initial_client_secret
  #   }
  # }

  # --- Okta variant (needs a DCR management/SSWS token from the Okta org) ---
  # okta = {
  #   name          = "my-dcr-provider-okta"
  #   display_name  = "My DCR Provider (Okta)"
  #   issuer        = "https://my-org.okta.com/oauth2/default"
  #
  #   dcr_config = {
  #     dcr_token = var.okta_dcr_token
  #   }
  # }

  # --- HTTP variant (a custom DCR server reachable from Konnect) ---
  # http = {
  #   name          = "my-dcr-provider-http"
  #   display_name  = "My DCR Provider (HTTP)"
  #   issuer        = "https://issuer.example.com"
  #
  #   dcr_config = {
  #     dcr_base_url                = "https://dcr.example.com"
  #     api_key                     = var.http_dcr_api_key
  #     allow_multiple_credentials  = false
  #     disable_event_hooks         = false
  #     disable_refresh_secret      = false
  #   }
  # }
}

resource "konnect_application_auth_strategy" "my_authstrategy" {
  openid_connect = {
    strategy_type = "openid_connect"
    name          = "my-auth-strategy"
    display_name  = "My Auth Strategy (DCR)"

    dcr_provider_id = konnect_dcr_provider.my_dcrprovider.id

    labels = {
      team = "platform"
    }

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
