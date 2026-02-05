resource "konnect_gateway_plugin_ai_prompt_compressor" "my_gatewaypluginaipromptcompressor" {
  config = {
    compression_ranges = [
      {
        max_tokens = 8
        min_tokens = 8
        value      = 5.24
      }
    ]
    compressor_type   = "target_token"
    compressor_url    = "...my_compressor_url..."
    keepalive_timeout = 6.37
    log_text_data     = true
    message_type = [
      "user"
    ]
    stop_on_error = true
    timeout       = 1.73
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
  updated_at = 7
}