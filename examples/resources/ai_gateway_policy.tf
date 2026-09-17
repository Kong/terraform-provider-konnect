resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "My AI Gateway"
  name         = "my-ai-gateway"
}

resource "konnect_ai_gateway_policy" "my_aigatewaypolicy" {
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name         = "pii-sanitizer"
  display_name = "PII Sanitizer Policy"
  type         = "ai-sanitizer"
  config = jsonencode({
    "sanitization_mode" = "INPUT"
    "anonymize" = [
      "all_and_credentials",
    ]
    "redact_type"       = "placeholder"
    "recover_redacted"  = true
    "block_if_detected" = false
    "stop_on_error"     = true
  })
}
