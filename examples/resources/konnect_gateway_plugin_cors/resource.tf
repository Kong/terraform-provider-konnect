resource "konnect_gateway_plugin_cors" "my_gatewayplugincors" {
  config = {
    allow_origin_absent = true
    credentials         = true
    exposed_headers = [
      "..."
    ]
    headers = [
      "..."
    ]
    max_age = 8.8
    methods = [
      "DELETE"
    ]
    origins = [
      "..."
    ]
    preflight_continue = true
    private_network    = true
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 3
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
  updated_at = 2
}