resource "konnect_gateway_plugin_oauth2_introspection" "my_gatewaypluginoauth2introspection" {
  config = {
    anonymous           = "...my_anonymous..."
    authorization_value = "...my_authorization_value..."
    consumer_by         = "username"
    custom_claims_forward = [
      "..."
    ]
    custom_introspection_headers = {
      key = "value"
    }
    hide_credentials   = true
    introspect_request = false
    introspection_url  = "...my_introspection_url..."
    keepalive          = 2
    run_on_preflight   = true
    timeout            = 8
    token_type_hint    = "...my_token_type_hint..."
    ttl                = 1.29
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 3
  enabled          = false
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
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
  protocols = [
    "grpc"
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
  updated_at = 8
}