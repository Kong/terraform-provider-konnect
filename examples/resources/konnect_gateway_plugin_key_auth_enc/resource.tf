resource "konnect_gateway_plugin_key_auth_enc" "my_gatewaypluginkeyauthenc" {
  config = {
    anonymous        = "...my_anonymous..."
    hide_credentials = false
    key_in_body      = false
    key_in_header    = false
    key_in_query     = false
    key_names = [
      "..."
    ]
    realm            = "...my_realm..."
    run_on_preflight = true
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
    "tcp"
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