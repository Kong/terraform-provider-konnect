resource "konnect_gateway_plugin_ai_prompt_template" "my_gatewaypluginaiprompttemplate" {
  config = {
    allow_untemplated_requests = true
    log_original_request       = false
    max_request_body_size      = 9
    templates = [
      {
        name     = "...my_name..."
        template = "...my_template..."
      }
    ]
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