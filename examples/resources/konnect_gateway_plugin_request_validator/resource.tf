resource "konnect_gateway_plugin_request_validator" "my_gatewaypluginrequestvalidator" {
  config = {
    allowed_content_types = [
      "..."
    ]
    body_schema                       = "...my_body_schema..."
    content_type_parameter_validation = false
    parameter_schema = [
      {
        explode  = true
        in       = "query"
        name     = "...my_name..."
        required = false
        schema   = "...my_schema..."
        style    = "matrix"
      }
    ]
    verbose_response = true
    version          = "draft6"
  }
  consumer = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 5
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
    "grpcs"
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