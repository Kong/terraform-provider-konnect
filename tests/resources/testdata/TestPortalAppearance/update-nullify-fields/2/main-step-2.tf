resource "konnect_portal" "my_portal" {
  name                      = "My New Portal for appearance nullify"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "portal.example-nullify-appearance.com"
  is_public                 = false
  rbac_enabled              = false
}

resource "konnect_portal_appearance" "nullify_appearance_test" {
  portal_id = konnect_portal.my_portal.id
}
