resource "konnect_gateway_plugin_request_termination" "my_gatewaypluginrequesttermination" {
  config = {
    body         = "...my_body..."
    content_type = "...my_content_type..."
    echo         = true
    message      = "...my_message..."
    status_code  = 536
    trigger      = "...my_trigger..."
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