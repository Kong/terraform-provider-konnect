resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI AWS Guardrails Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_aws_guardrails" "my_plugin" {
  enabled = true

  config = {
    aws_region           = "us-west-2"
    guardrails_id        = "gr1234567890"
    guardrails_version   = "1"
    guarding_mode        = "BOTH"
    text_source          = "concatenate_user_content"
    response_buffer_size = 200
    timeout              = 20000
    allow_masking        = true
    log_blocked_content  = true
    ssl_verify           = false
    stop_on_error        = false
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
