resource "konnect_gateway_plugin_loggly" "my_gatewaypluginloggly" {
  config = {
    client_errors_severity = "notice"
    custom_fields_by_lua = {
      key = "value"
    }
    host                   = "...my_host..."
    key                    = "...my_key..."
    log_level              = "alert"
    port                   = 51730
    server_errors_severity = "crit"
    successful_severity    = "alert"
    tags = [
      "..."
    ]
    timeout = 9.55
  }
  consumer = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 8
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
    "http"
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
  updated_at = 10
}