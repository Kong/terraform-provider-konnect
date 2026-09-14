resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - identity"
  name         = "tf-test-aigw-dp-cert"
}

resource "konnect_ai_gateway_identity_provider" "my_aigatewayidentityprovider" {
    provider = konnect-beta
    gateway_id = konnect_ai_gateway.my_aigateway.id
    key_auth = {
        config = jsonencode({
            key_in_header         = true
        })
        display_name = "Okta AI SE"
        
        name = "tf-test-key-auth-identity-provider"
    }

    // config is free form, and the API can add more keys to it, so we ignore changes to it to avoid unnecessary diff in tests.
    lifecycle {
        ignore_changes = [
            key_auth.config
        ]
    }
}

resource "konnect_ai_gateway_identity_provider" "my_aigatewayidentityprovider_oidc" {
  provider = konnect-beta
  gateway_id = konnect_ai_gateway.my_aigateway.id
  openid_connect = {
    config = jsonencode({
      cache_tokens_salt = "my_cache_tokens_salt"
      issuer                   = "https://dev-123456.example.com"
      ssl_verify = true
    })
    display_name = "TF Test OpenID Connect"

    name = "tf-test-openid-connect-identity-provider"
  }

// config is free form, and the API can add more keys to it, so we ignore changes to it to avoid unnecessary diff in tests.
  lifecycle {
        ignore_changes = [
            openid_connect.config
        ]
    }
}