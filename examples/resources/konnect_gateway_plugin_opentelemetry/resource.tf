resource "konnect_gateway_plugin_opentelemetry" "my_gatewaypluginopentelemetry" {
  config = {
    access_logs_endpoint = "...my_access_logs_endpoint..."
    batch_flush_delay    = 7
    batch_span_count     = 5
    connect_timeout      = 1207240418
    header_type          = "gcp"
    headers = {
      key = "value"
    }
    http_response_header_for_traceid = "...my_http_response_header_for_traceid..."
    logs_endpoint                    = "...my_logs_endpoint..."
    metrics = {
      enable_bandwidth_metrics       = true
      enable_consumer_attribute      = false
      enable_latency_metrics         = false
      enable_request_metrics         = true
      enable_upstream_health_metrics = false
      endpoint                       = "...my_endpoint..."
      push_interval                  = 0.85
    }
    propagation = {
      clear = [
        "..."
      ]
      default_format = "aws"
      extract = [
        "instana"
      ]
      inject = [
        "b3"
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
      key = "value"
    }
    sampling_rate     = 0.37
    sampling_strategy = "parent_drop_probability_fallback"
    send_timeout      = 1637096441
    traces_endpoint   = "...my_traces_endpoint..."
  }
  consumer = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 0
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
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
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
  updated_at = 8
}