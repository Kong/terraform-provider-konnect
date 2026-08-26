resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI LLM As Judge Plugin CRUD"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_llm_as_judge" "my_plugin" {
  enabled = true

  config = {
    prompt                   = "You are a strict evaluator. Score the response between 1 and 100."
    http_timeout             = 60000
    https_verify             = true
    ignore_assistant_prompts = true
    ignore_system_prompts    = true
    ignore_tool_prompts      = true
    message_countback        = 5
    sampling_rate            = 1

    llm = {
      route_type = "llm/v1/chat"

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
    }
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
