---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_audit_log Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  AuditLog Resource
---

# konnect_audit_log (Resource)

AuditLog Resource

## Example Usage

```terraform
resource "konnect_audit_log" "my_auditlog" {
  authorization         = "Bearer sometoken"
  enabled               = true
  endpoint              = "https://example.com/audit-logs"
  log_format            = "cps"
  skip_ssl_verification = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `authorization` (String) The value to include in the `Authorization` header when sending audit logs to the webhook.
- `enabled` (Boolean) Indicates if the data should be sent to the webhook.
- `endpoint` (String) The endpoint that will receive audit log messages.
- `log_format` (String) The output format of each log messages. Default: "cef"; must be one of ["cef", "json", "cps"]
- `skip_ssl_verification` (Boolean) Indicates if the SSL certificate verification of the host endpoint should be skipped when delivering payloads.
We strongly recommend not setting this to 'true' as you are subject to man-in-the-middle and other attacks.
This option should be considered only for self-signed SSL certificates used in a non-production environment.

### Read-Only

- `updated_at` (String) Timestamp when this webhook was last updated. Initial value is 0001-01-01T00:00:0Z.
