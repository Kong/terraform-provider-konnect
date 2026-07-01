resource "konnect_portal" "my_portal" {
  force_destroy = "true"

  name                   = "My v3 portal for team"
  authentication_enabled = true
  rbac_enabled           = true
  default_api_visibility = "public"
}

resource "konnect_portal_team" "my_portal_team" {
  portal_id            = konnect_portal.my_portal.id
  name                 = "IDM - Developers"
  description          = "Identity Management team"
  can_own_applications = true
}
