resource "konnect_gateway_plugin_mocking" "my_gatewaypluginmocking" {
  config = {
    api_specification          = "...my_api_specification..."
    api_specification_filename = "...my_api_specification_filename..."
    custom_base_path           = "...my_custom_base_path..."
    include_base_path          = true
    included_status_codes = [
      5
    ]
    max_delay_time     = 1.21
    min_delay_time     = 6.51
    random_delay       = true
    random_examples    = false
    random_status_code = true
  }
  consumer = {
    id = "...my_id..."
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