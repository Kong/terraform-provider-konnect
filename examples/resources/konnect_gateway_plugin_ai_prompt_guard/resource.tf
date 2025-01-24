resource "konnect_gateway_plugin_ai_prompt_guard" "my_gatewaypluginaipromptguard" {
  config = {
    allow_all_conversation_history = false
    allow_patterns = [
      "..."
    ]
    deny_patterns = [
      "..."
    ]
    match_all_roles       = false
    max_request_body_size = 10
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
    "tls_passthrough"
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