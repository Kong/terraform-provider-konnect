resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - identity"
  name         = "tf-test-aigw-dp-cert"
}

resource "konnect_ai_gateway_identity_provider" "my_aigatewayidentityprovider" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  key_auth = {
    config = jsonencode({
      key_in_header = true
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
