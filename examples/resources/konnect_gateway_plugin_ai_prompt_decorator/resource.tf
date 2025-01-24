resource "konnect_gateway_plugin_ai_prompt_decorator" "my_gatewaypluginaipromptdecorator" {
  config = {
    max_request_body_size = 3
    prompts = {
      append = [
        {
          content = "...my_content..."
          role    = "user"
        }
      ]
      prepend = [
        {
          content = "...my_content..."
          role    = "system"
        }
      ]
    }
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
    "tls"
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