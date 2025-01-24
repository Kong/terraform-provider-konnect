resource "konnect_gateway_plugin_http_log" "my_gatewaypluginhttplog" {
  config = {
    content_type = "application/json; charset=utf-8"
    custom_fields_by_lua = {
      key = jsonencode("value"),
    }
    flush_timeout = 8.97
    headers = {
      key = jsonencode("value"),
    }
    http_endpoint = "...my_http_endpoint..."
    keepalive     = 2.25
    method        = "PUT"
    queue = {
      concurrency_limit    = 0
      initial_retry_delay  = 154435.3
      max_batch_size       = 297113
      max_bytes            = 7
      max_coalescing_delay = 198.88
      max_entries          = 603921
      max_retry_delay      = 290740.88
      max_retry_time       = 4.99
    }
    queue_size  = 3
    retry_count = 10
    timeout     = 6.49
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