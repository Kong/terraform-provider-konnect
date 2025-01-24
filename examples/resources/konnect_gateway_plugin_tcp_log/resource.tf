resource "konnect_gateway_plugin_tcp_log" "my_gatewayplugintcplog" {
  config = {
    custom_fields_by_lua = {
      key = jsonencode("value"),
    }
    host      = "...my_host..."
    keepalive = 1.57
    port      = 54956
    timeout   = 5.53
    tls       = false
    tls_sni   = "...my_tls_sni..."
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