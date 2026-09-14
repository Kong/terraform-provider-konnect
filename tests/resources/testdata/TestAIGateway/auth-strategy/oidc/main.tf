resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - auth-strategy"
  name         = "tf-test-aigw-auth-strategy"
}

resource "konnect_ai_gateway_auth_strategy" "my_aigatewayauthstrategy_oidc" {
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
    ignore_changes = [openid_connect.config]
  }
}
