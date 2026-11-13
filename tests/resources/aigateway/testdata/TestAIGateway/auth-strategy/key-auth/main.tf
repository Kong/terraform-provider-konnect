resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - auth-strategy"
  name         = "tf-test-aigw-auth-strategy"
}

resource "konnect_ai_gateway_auth_strategy" "my_aigatewayauthstrategy" {
  gateway_id = konnect_ai_gateway.my_aigateway.id

  key_auth = {
    config = {
      hide_credentials = false
      key_in_header    = true
      key_in_query     = true
      key_in_body      = false
      key_names = [
        "apikey"
      ]
    }
    display_name = "TF Test Key Auth"
    name         = "tf-test-key-auth"
  }
}
