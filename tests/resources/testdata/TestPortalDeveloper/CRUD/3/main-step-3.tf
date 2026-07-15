resource "konnect_portal" "test_portal" {
  name                      = "test_portal_for_developers"
  auto_approve_developers   = false
  auto_approve_applications = false
}

resource "konnect_portal_developer" "test_developer" {
  portal_id             = konnect_portal.test_portal.id
  email                 = "developer@example.com"
  full_name             = "Test Developer"
  send_invitation_email = false
  status                = "pending"
}

