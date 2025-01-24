resource "konnect_gateway_plugin_loggly" "my_gatewaypluginloggly" {
  config = {
    client_errors_severity = "alert"
    custom_fields_by_lua = {
      key = jsonencode("value"),
    }
    host                   = "...my_host..."
    key                    = "...my_key..."
    log_level              = "debug"
    port                   = 51730
    server_errors_severity = "info"
    successful_severity    = "debug"
    tags = [
      "..."
    ]
    timeout = 9.55
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
}