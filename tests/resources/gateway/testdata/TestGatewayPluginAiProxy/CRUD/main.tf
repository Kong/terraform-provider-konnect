resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI Proxy Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_proxy" "my_plugin" {
  enabled = true

  config = {
    genai_category = "text/generation"
    llm_format     = "openai"
    route_type     = "llm/v1/chat"

    model = {
      name     = "claude-3-5-sonnet-20241022"
      provider = "anthropic"

      options = {
        anthropic_version = "2023-06-01"
        max_tokens        = 1024

        input_cost       = 3.00
        output_cost      = 15.00
        cache_read_cost  = 0.30
        cache_write_cost = 3.75
      }
    }
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
