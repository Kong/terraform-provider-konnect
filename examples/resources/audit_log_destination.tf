resource "konnect_audit_log_destination" "my_auditlogdestination" {
  authorization         = "Bearer sometoken"
  endpoint              = "https://example.com/audit-logs"
  log_format            = "cef"
  name                  = "My Destination"
  skip_ssl_verification = false
}