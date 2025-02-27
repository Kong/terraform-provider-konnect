resource "konnect_gateway_plugin_injection_protection" "my_gatewayplugininjectionprotection" {
  config = {
    custom_injections = [
      {
        name  = "...my_name..."
        regex = "...my_regex..."
      }
    ]
    enforcement_mode  = "block"
    error_message     = "...my_error_message..."
    error_status_code = 416
    injection_types = [
      "sql"
    ]
    locations = [
      "path_and_query"
    ]
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
}