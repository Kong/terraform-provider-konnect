resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI Proxy Advanced Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_proxy_advanced" "my_plugin" {
  enabled = true

  config = {
    genai_category = "text/generation"
    llm_format     = "openai"

    targets = [
      {
        description = "anthropic target updated"
        route_type  = "llm/v1/chat"
        weight      = 50

        model = {
          name     = "claude-3-5-sonnet-20241022"
          provider = "anthropic"

          options = {
            anthropic_version = "2023-06-01"
            max_tokens        = 2048

            input_cost       = 5.00
            output_cost      = 20.00
            cache_read_cost  = 0.50
            cache_write_cost = 4.75
          }
        }
      },
    ]
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
