resource "konnect_gateway_plugin_oas_validation" "my_gatewaypluginoasvalidation" {
  config = {
    allowed_header_parameters                    = "...my_allowed_header_parameters..."
    api_spec                                     = "...my_api_spec..."
    api_spec_encoded                             = false
    custom_base_path                             = "...my_custom_base_path..."
    header_parameter_check                       = false
    include_base_path                            = true
    notify_only_request_validation_failure       = false
    notify_only_response_body_validation_failure = true
    query_parameter_check                        = true
    validate_request_body                        = true
    validate_request_header_params               = true
    validate_request_query_params                = false
    validate_request_uri_params                  = true
    validate_response_body                       = false
    verbose_response                             = true
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
    "wss"
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