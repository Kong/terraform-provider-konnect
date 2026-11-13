resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI AWS Guardrails Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_aws_guardrails" "my_plugin" {
  enabled = true

  config = {
    aws_region           = "us-east-1"
    guardrails_id        = "gr1234567890"
    guardrails_version   = "DRAFT"
    guarding_mode        = "INPUT"
    text_source          = "concatenate_all_content"
    response_buffer_size = 100
    timeout              = 10000
    allow_masking        = false
    log_blocked_content  = false
    ssl_verify           = false
    stop_on_error        = true
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
