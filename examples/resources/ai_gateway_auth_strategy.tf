resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "My AI Gateway"
  name         = "my-ai-gateway"
}

resource "konnect_ai_gateway_auth_strategy" "my_aigatewayauthstrategy" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  key_auth = {
    name         = "key-auth"
    display_name = "Key Auth"
    config = {
      key_names     = ["apikey"]
      key_in_header = true
      key_in_query  = true
      key_in_body   = false
    }
  }
}
