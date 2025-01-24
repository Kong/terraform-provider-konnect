resource "konnect_gateway_plugin_opa" "my_gatewaypluginopa" {
  config = {
    include_body_in_opa_input             = true
    include_consumer_in_opa_input         = true
    include_parsed_json_body_in_opa_input = true
    include_route_in_opa_input            = false
    include_service_in_opa_input          = true
    include_uri_captures_in_opa_input     = false
    opa_host                              = "...my_opa_host..."
    opa_path                              = "...my_opa_path..."
    opa_port                              = 43549
    opa_protocol                          = "https"
    ssl_verify                            = false
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
    "http"
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