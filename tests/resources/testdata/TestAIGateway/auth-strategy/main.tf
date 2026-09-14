resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - auth-strategy"
  name         = "tf-test-aigw-auth-strategy"
}

resource "konnect_ai_gateway_auth_strategy" "my_aigatewayauthstrategy" {
  provider   = konnect-beta
  gateway_id = konnect_ai_gateway.my_aigateway.id

  key_auth = {
    config = jsonencode({
      hide_credentials = false
      key_in_header    = true
      key_in_query     = true
      key_in_body = false
      key_names = [
        "apikey"
      ]
    })
    display_name = "TF Test Key Auth"
    name         = "tf-test-key-auth"
  }

  lifecycle {
    # free form, ignore to avoid test failure when new fields supported by server
    ignore_changes = [ key_auth.config ]
  }
}

resource "konnect_ai_gateway_auth_strategy" "my_aigatewayauthstrategy_oidc" {
  provider   = konnect-beta
  gateway_id = konnect_ai_gateway.my_aigateway.id

  openid_connect = {
    config = jsonencode({
      cache_tokens_salt = "my_cache_tokens_salt"
      client_alg = [
        "RS256"
      ]
      client_auth = [
        "client_secret_basic"
      ]
      client_id = [
        "test-client-id"
      ]
      client_secret = [
        "test-client-secret"
      ]
      issuer           = "https://example.com"
      jwks_endpoint    = "https://example.com/.well-known/jwks.json"
      hide_credentials = false
    })
    display_name = "TF Test OpenID Connect"
    name         = "tf-test-oidc"
  }

  lifecycle {
    # free form, ignore to avoid test failure when new fields supported by server
    ignore_changes = [ openid_connect.config ]
  }
}
