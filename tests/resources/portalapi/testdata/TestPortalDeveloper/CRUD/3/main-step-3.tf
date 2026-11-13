resource "konnect_portal" "test_portal" {
  name                      = "test_portal_for_developers"
  auto_approve_developers   = false
  auto_approve_applications = false
}

resource "konnect_portal_team" "my_team" {
  portal_id = konnect_portal.test_portal.id
  name        = "TFAcceptancePortalDeveloperTeam"
  description = "TF acceptance test portal team for adding developer."
}

resource "konnect_portal_developer" "test_developer" {
  portal_id             = konnect_portal.test_portal.id
  email                 = "developer@example.com"
  full_name             = "Test Developer"
  send_invitation_email = false
  status                = "pending"
}

resource "konnect_portal_team_developer" "my_portalteamdeveloper" {
  developer_id = konnect_portal_developer.test_developer.id
  portal_id    = konnect_portal.test_portal.id
  team_id      = konnect_portal_team.my_team.id
}
