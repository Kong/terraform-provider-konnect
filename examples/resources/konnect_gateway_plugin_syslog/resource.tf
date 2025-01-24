resource "konnect_gateway_plugin_syslog" "my_gatewaypluginsyslog" {
  config = {
    client_errors_severity = "emerg"
    custom_fields_by_lua = {
      key = jsonencode("value"),
    }
    facility               = "local7"
    log_level              = "info"
    server_errors_severity = "alert"
    successful_severity    = "notice"
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