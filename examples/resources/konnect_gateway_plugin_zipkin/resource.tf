resource "konnect_gateway_plugin_zipkin" "my_gatewaypluginzipkin" {
  config = {
    connect_timeout                  = 1674102719
    default_header_type              = "w3c"
    default_service_name             = "...my_default_service_name..."
    header_type                      = "jaeger"
    http_endpoint                    = "...my_http_endpoint..."
    http_response_header_for_traceid = "...my_http_response_header_for_traceid..."
    http_span_name                   = "method_path"
    include_credential               = false
    local_service_name               = "...my_local_service_name..."
    phase_duration_flavor            = "annotations"
    propagation = {
      clear = [
        "..."
      ]
      default_format = "b3-single"
      extract = [
        "b3"
      ]
      inject = [
        "w3c"
      ]
    }
    queue = {
      concurrency_limit    = 1
      initial_retry_delay  = 489129.82
      max_batch_size       = 540049
      max_bytes            = 6
      max_coalescing_delay = 2815.46
      max_entries          = 977371
      max_retry_delay      = 76225.98
      max_retry_time       = 2.97
    }
    read_timeout = 1936940031
    sample_ratio = 0.6
    send_timeout = 2045499472
    static_tags = [
      {
        name  = "...my_name..."
        value = "...my_value..."
      }
    ]
    tags_header        = "...my_tags_header..."
    traceid_byte_count = 8
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