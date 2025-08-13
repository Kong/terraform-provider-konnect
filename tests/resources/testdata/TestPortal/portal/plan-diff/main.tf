resource "konnect_portal" "my_portal" {
  name                      = "My v2 Developer Portal"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "tf-example-domain-portal-test.example.com"
  is_public                 = false
  rbac_enabled              = false
}