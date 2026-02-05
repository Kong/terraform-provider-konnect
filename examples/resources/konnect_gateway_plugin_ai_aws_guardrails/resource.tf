resource "konnect_gateway_plugin_ai_aws_guardrails" "my_gatewaypluginaiawsguardrails" {
  config = {
    allow_masking         = true
    aws_access_key_id     = "...my_aws_access_key_id..."
    aws_assume_role_arn   = "...my_aws_assume_role_arn..."
    aws_region            = "...my_aws_region..."
    aws_role_session_name = "...my_aws_role_session_name..."
    aws_secret_access_key = "...my_aws_secret_access_key..."
    aws_sts_endpoint_url  = "...my_aws_sts_endpoint_url..."
    guarding_mode         = "INPUT"
    guardrails_id         = "...my_guardrails_id..."
    guardrails_version    = "...my_guardrails_version..."
    response_buffer_size  = 4.85
    ssl_verify            = true
    stop_on_error         = false
    text_source           = "concatenate_all_content"
    timeout               = 4.59
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 4
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
    "grpc"
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
  updated_at = 9
}