# Create a new Dev Portal
resource "konnect_portal" "my_portal" {
  name                      = "My New Portal"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "demo.example.com"
  is_public                 = false
  rbac_enabled              = false
}

# Create a new team in the Dev Portal
resource "konnect_portal_team" "myteam" {
    portal_id = konnect_portal.my_portal.id
    name = "My platform team"
    description = "Description for my platform team"
}