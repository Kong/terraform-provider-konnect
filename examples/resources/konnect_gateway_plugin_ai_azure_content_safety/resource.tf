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
    guarding_mode         = "INPUT"
    halt_on_blocklist_hit = false
    output_type           = "FourSeverityLevels"
    response_buffer_size  = 7.56
    reveal_failure_reason = false
    ssl_verify            = false
    stop_on_error         = false
    text_source           = "concatenate_user_content"
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 10
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
  updated_at = 10
}