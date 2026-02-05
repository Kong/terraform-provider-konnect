resource "konnect_audit_log" "my_auditlog" {
  authorization         = "Bearer sometoken"
  enabled               = true
  endpoint              = "https://example.com/audit-logs"
  log_format            = "cps"
  skip_ssl_verification = false
}