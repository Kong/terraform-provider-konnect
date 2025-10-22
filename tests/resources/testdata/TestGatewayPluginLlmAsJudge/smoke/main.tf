resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For LLM As Judge Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_llm_as_judge" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin"

  prompt                   = <<EOF
You are a strict evaluator. You will be given a request and a response.
Your task is to judge whether the response is correct or incorrect. You must
assign a score between 1 and 100, where: 100 represents a completely correct
and ideal response, 1 represents a completely incorrect or irrelevant response.
Your score must be a single number only — no text, labels, or explanations.
Use the full range of values (e.g., 13, 47, 86), not just round numbers like
10, 50, or 100. Be accurate and consistent, as this score will be used by another
model for learning and evaluation.
EOF
  http_timeout             = 60000
  https_verify             = true
  ignore_assistant_prompts = true
  ignore_system_prompts    = true
  ignore_tool_prompts      = true
  sampling_rate            = 1

  llm = {

    auth = {
      allow_override = false
      header_name    = "Authorization"
      header_value   = "Bearer var.openai_api_key"
    }

    logging = {
      log_payloads   = true
      log_statistics = true
    }

    model = {
      name     = "gpt-4o"
      provider = "openai"

      options = {
        temperature = 2
        max_tokens  = 5
        top_p       = 1

        cohere = {
          embedding_input_type = "classification"
        }
      }
    }
    route_type = "llm/v1/chat"
  }
  message_countback = 3

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
