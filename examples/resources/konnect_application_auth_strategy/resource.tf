resource "konnect_application_auth_strategy" "my_applicationauthstrategy" {
  key_auth = {
    configs = {
      key_auth = {
        key_names = [
          "..."
        ]
      }
    }
    display_name = "...my_display_name..."
    labels = {
      key = "value",
    }
    name          = "...my_name..."
    strategy_type = "key_auth"
  }
  openid_connect = {
    configs = {
      openid_connect = {
        additional_properties = "{ \"see\": \"documentation\" }"
        auth_methods = [
          "..."
        ]
        credential_claim = [
          "..."
        ]
        issuer = "...my_issuer..."
        labels = {
          key = "value",
        }
        scopes = [
          "..."
        ]
      }
    }
    dcr_provider_id = "184e9c55-484e-4f4f-9de7-f6001d8ab0e7"
    display_name    = "...my_display_name..."
    labels = {
      key = "value",
    }
    name          = "...my_name..."
    strategy_type = "openid_connect"
  }
}