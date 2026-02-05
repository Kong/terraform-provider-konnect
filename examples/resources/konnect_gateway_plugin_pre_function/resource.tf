resource "konnect_gateway_plugin_pre_function" "my_gatewaypluginprefunction" {
  config = {
    access = [
      "..."
    ]
    body_filter = [
      "..."
    ]
    certificate = [
      "..."
    ]
    header_filter = [
      "..."
    ]
    log = [
      "..."
    ]
    rewrite = [
      "..."
    ]
    ws_client_frame = [
      "..."
    ]
    ws_close = [
      "..."
    ]
    ws_handshake = [
      "..."
    ]
    ws_upstream_frame = [
      "..."
    ]
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 6
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
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
  protocols = [
    "udp"
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
  updated_at = 3
}