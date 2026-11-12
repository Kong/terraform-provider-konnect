resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI Azure Content Safety Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_azure_content_safety" "my_plugin" {
  enabled = true

  config = {
    content_safety_url    = "https://my-content-safety.cognitiveservices.azure.com"
    azure_api_version     = "2023-10-01"
    guarding_mode         = "BOTH"
    output_type           = "EightSeverityLevels"
    text_source           = "concatenate_user_content"
    response_buffer_size  = 200
    halt_on_blocklist_hit = true
    reveal_failure_reason = false
    log_blocked_content   = false
    ssl_verify            = false
    stop_on_error         = true

    blocklist_names = [
      "blocklist-two",
    ]

    categories = [
      {
        name            = "Hate"
        rejection_level = 6
      },
      {
        name            = "SelfHarm"
        rejection_level = 4
      },
    ]
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
