resource "konnect_portal_classic" "my_portal" {
  name                      = "My v2 Developer Portal for team"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "tf-example-domain-portal-team.example.com"
  is_public                 = false
  rbac_enabled              = false
}

resource "konnect_portal_team" "my_test_team" {
    portal_id = konnect_portal_classic.my_portal.id
    name = "My test team"
    description = "Description for my test team"
}