resource "konnect_gateway_plugin_opentelemetry" "my_gatewaypluginopentelemetry" {
  config = {
    batch_flush_delay = 7
    batch_span_count  = 5
    connect_timeout   = 1207240418
    header_type       = "w3c"
    headers = {
      key = jsonencode("value"),
    }
    http_response_header_for_traceid = "...my_http_response_header_for_traceid..."
    logs_endpoint                    = "...my_logs_endpoint..."
    propagation = {
      clear = [
        "..."
      ]
      default_format = "w3c"
      extract = [
        "jaeger"
      ]
      inject = [
        "w3c"
      ]
    }
    queue = {
      concurrency_limit    = 2
      initial_retry_delay  = 226722.01
      max_batch_size       = 779071
      max_bytes            = 9
      max_coalescing_delay = 763.01
      max_entries          = 975127
      max_retry_delay      = 892016.33
      max_retry_time       = 0.21
    }
    read_timeout = 1485093464
    resource_attributes = {
      key = jsonencode("value"),
    }
    sampling_rate   = 0.37
    send_timeout    = 1637096441
    traces_endpoint = "...my_traces_endpoint..."
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