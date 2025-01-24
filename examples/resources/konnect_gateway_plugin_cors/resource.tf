resource "konnect_gateway_plugin_cors" "my_gatewayplugincors" {
  config = {
    credentials = true
    exposed_headers = [
      "..."
    ]
    headers = [
      "..."
    ]
    max_age = 8.8
    methods = [
      "HEAD"
    ]
    origins = [
      "..."
    ]
    preflight_continue = true
    private_network    = true
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