resource "konnect_portal_custom_domain" "my_portalcustomdomain" {
  enabled   = false
  hostname  = "...my_hostname..."
  portal_id = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  ssl = {
    custom_certificate         = "...my_custom_certificate..."
    custom_private_key         = "...my_custom_private_key..."
    domain_verification_method = "custom_certificate"
    skip_ca_check              = true
  }
}