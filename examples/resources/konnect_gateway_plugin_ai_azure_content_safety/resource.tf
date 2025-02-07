resource "konnect_gateway_plugin_ai_azure_content_safety" "my_gatewaypluginaiazurecontentsafety" {
  config = {
    azure_api_version          = "...my_azure_api_version..."
    azure_client_id            = "...my_azure_client_id..."
    azure_client_secret        = "...my_azure_client_secret..."
    azure_tenant_id            = "...my_azure_tenant_id..."
    azure_use_managed_identity = true
    blocklist_names = [
      "..."
    ]
    categories = [
      {
        name            = "...my_name..."
        rejection_level = 9
      }
    ]
    content_safety_key    = "...my_content_safety_key..."
    content_safety_url    = "...my_content_safety_url..."
    halt_on_blocklist_hit = false
    output_type           = "FourSeverityLevels"
    reveal_failure_reason = false
    text_source           = "concatenate_user_content"
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