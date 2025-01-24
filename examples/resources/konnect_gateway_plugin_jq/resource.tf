resource "konnect_gateway_plugin_jq" "my_gatewaypluginjq" {
  config = {
    request_if_media_type = [
      "..."
    ]
    request_jq_program = "...my_request_jq_program..."
    request_jq_program_options = {
      ascii_output   = false
      compact_output = false
      join_output    = true
      raw_output     = true
      sort_keys      = false
    }
    response_if_media_type = [
      "..."
    ]
    response_if_status_code = [
      188
    ]
    response_jq_program = "...my_response_jq_program..."
    response_jq_program_options = {
      ascii_output   = false
      compact_output = true
      join_output    = true
      raw_output     = false
      sort_keys      = false
    }
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