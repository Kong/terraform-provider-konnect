resource "konnect_gateway_plugin_oauth2_introspection" "my_gatewaypluginoauth2introspection" {
  config = {
    anonymous           = "...my_anonymous..."
    authorization_value = "...my_authorization_value..."
    consumer_by         = "client_id"
    custom_claims_forward = [
      "..."
    ]
    custom_introspection_headers = {
      key = jsonencode("value"),
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
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
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
  protocols = [
    "grpcs"
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