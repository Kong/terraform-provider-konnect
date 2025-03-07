resource "konnect_gateway_plugin_json_threat_protection" "my_gatewaypluginjsonthreatprotection" {
  config = {
    enforcement_mode             = "log_only"
    error_message                = "...my_error_message..."
    error_status_code            = 434
    max_array_element_count      = 96675195
    max_body_size                = 99378519
    max_container_depth          = 2031079601
    max_object_entry_count       = 916870322
    max_object_entry_name_length = 1383216872
    max_string_value_length      = 1149380350
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