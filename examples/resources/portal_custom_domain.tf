resource "konnect_portal_custom_domain" "my_portal_custom_domain" {
  enabled   = false
  hostname  = "my.custom.domain"
  portal_id = konnect_portal.my_portal.id
  ssl = {
    domain_verification_method = "http"
  }
}