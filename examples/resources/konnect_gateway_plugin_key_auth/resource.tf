resource "konnect_gateway_plugin_key_auth" "my_gatewaypluginkeyauth" {
  config = {
    anonymous        = "...my_anonymous..."
    hide_credentials = false
    key_in_body      = false
    key_in_header    = false
    key_in_query     = true
    key_names = [
      "..."
    ]
    realm            = "...my_realm..."
    run_on_preflight = false
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
    "https"
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