resource "konnect_portal" "my_portal_for_custom_domain" {
  force_destroy = "true"
  name         = "My v3 portal name"
}

resource "konnect_portal_custom_domain" "my_portal_custom_domain" {
  enabled   = false
  hostname  = "my.tftestingorgcustom.domain"
  portal_id = konnect_portal.my_portal_for_custom_domain.id
  ssl = {
    domain_verification_method = "http"
  }
}