resource "konnect_gateway_plugin_datadog_tracing" "my_gatewayplugindatadogtracing" {
  config = {
    batch_flush_delay = 0
    batch_span_count  = 3
    connect_timeout   = 964301844
    endpoint          = "...my_endpoint..."
    environment       = "...my_environment..."
    read_timeout      = 582957688
    send_timeout      = 2117571111
    service_name      = "...my_service_name..."
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