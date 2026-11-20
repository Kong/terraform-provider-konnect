resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI Response Transformer Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_response_transformer" "my_plugin" {
  enabled = true

  config = {
    prompt                               = "Convert the upstream response body into a valid XML payload."
    http_timeout                         = 30000
    https_verify                         = true
    parse_llm_response_json_instructions = false

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
