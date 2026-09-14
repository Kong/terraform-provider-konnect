resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - policy"
  name         = "tf-test-aigw-policy"
}

resource "konnect_ai_gateway_policy" "my_aigatewaypolicy" {
  provider = konnect-beta
  config = jsonencode({
    "allow_all_conversation_history" = true
    "anonymize" = [
        "all_and_credentials",
      ]
    
    "block_if_detected" = false
    "custom_patterns"   = null
    "host"              = "localhost"
    "keepalive_timeout" = 60000
    "port"              = 8080
    "proxy_config" = {
        auth_password    = null
        auth_username    = null
        http_proxy_host  = null
        http_proxy_port  = null
        https_proxy_host = null
        https_proxy_port = null
        no_proxy         = null
        proxy_scheme     = "http"
      }
    "recover_redacted"             = true
    "redact_type"                  = "placeholder"
    "sanitization_mode"            = "INPUT"
    "scheme"                       = "http"
    "skip_logging_sanitized_items" = false
    "stop_on_error"                = true
    "timeout"                      = 10000
  })
  display_name = "My Cool AI PII Sanitizer Policy"
  gateway_id   = konnect_ai_gateway.my_aigateway.id

  name = "ai-pii-sanitizer-1234"
  type = "ai-sanitizer"

  lifecycle {
    ignore_changes = [
      config,
    ]
  }
}
