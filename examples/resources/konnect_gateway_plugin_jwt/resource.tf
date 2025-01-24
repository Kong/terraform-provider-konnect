resource "konnect_gateway_plugin_jwt" "my_gatewaypluginjwt" {
  config = {
    anonymous = "...my_anonymous..."
    claims_to_verify = [
      "exp"
    ]
    cookie_names = [
      "..."
    ]
    header_names = [
      "..."
    ]
    key_claim_name     = "...my_key_claim_name..."
    maximum_expiration = 24905725.8
    realm              = "...my_realm..."
    run_on_preflight   = true
    secret_is_base64   = false
    uri_param_names = [
      "..."
    ]
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = true
  id               = "...my_id..."
  instance_name    = "...my_instance_name..."
  ordering = {
    after = {
      access = [
        "..."
      ]
    }
    before = {
      access = [
        "..."
      ]
    }
  }
  protocols = [
    "wss"
  ]
  route = {
    id = "...my_id..."
  }
  service = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
}