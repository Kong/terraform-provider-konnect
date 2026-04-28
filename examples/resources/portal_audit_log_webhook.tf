resource "konnect_audit_log_destination" "my_auditlogdestination" {
authorization         = "Bearer sometoken"
endpoint              = "https://example.com/audit-logs"
log_format            = "cef"
name                  = "My Destination"
skip_ssl_verification = false
}

resource "konnect_portal" "my_portal" {
force_destroy          = "true"
default_api_visibility = "public"
name                   = "My Developer Portal"
}

resource "konnect_portal_audit_log_webhook" "my_portalauditlogwebhook" {
audit_log_destination_id = konnect_audit_log_destination.my_auditlogdestination.id
enabled                  = true
portal_id                = konnect_portal.my_portal.id
}
