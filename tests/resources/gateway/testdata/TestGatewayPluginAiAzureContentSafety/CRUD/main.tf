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
    guarding_mode         = "INPUT"
    output_type           = "FourSeverityLevels"
    text_source           = "concatenate_all_content"
    response_buffer_size  = 100
    halt_on_blocklist_hit = true
    reveal_failure_reason = true
    log_blocked_content   = false
    ssl_verify            = false
    stop_on_error         = true

    blocklist_names = [
      "blocklist-one",
    ]

    categories = [
      {
        name            = "Hate"
        rejection_level = 4
      },
      {
        name            = "Violence"
        rejection_level = 2
      },
    ]
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
