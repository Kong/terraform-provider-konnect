resource "konnect_gateway_plugin_hmac_auth" "my_gatewaypluginhmacauth" {
  config = {
    algorithms = [
      "hmac-sha256"
    ]
    anonymous  = "...my_anonymous..."
    clock_skew = 5.59
    enforce_headers = [
      "..."
    ]
    hide_credentials      = true
    realm                 = "...my_realm..."
    validate_request_body = true
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
    "tls_passthrough"
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