resource "konnect_gateway_plugin_ai_sanitizer" "my_gatewaypluginaisanitizer" {
  config = {
    anonymize = [
      "ssn"
    ]
    block_if_detected = false
    custom_patterns = [
      {
        name  = "...my_name..."
        regex = "...my_regex..."
        score = 0.99
      }
    ]
    host                         = "...my_host..."
    keepalive_timeout            = 7.7
    port                         = 6.01
    recover_redacted             = true
    redact_type                  = "placeholder"
    sanitization_mode            = "OUTPUT"
    scheme                       = "...my_scheme..."
    skip_logging_sanitized_items = false
    stop_on_error                = false
    timeout                      = 1.92
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 2
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
    "grpc"
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
  updated_at = 9
}