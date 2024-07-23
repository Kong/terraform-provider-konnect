resource "konnect_portal" "my_portal" {
  name                      = "My New Portal"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "demo.example.com"
  is_public                 = false
  rbac_enabled              = false
}
