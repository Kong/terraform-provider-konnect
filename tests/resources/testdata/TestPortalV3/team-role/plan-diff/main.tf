resource "konnect_portal" "my_portal" {
  force_destroy = "true"

  name                   = "My v3 portal for team role"
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

# The entity the role is scoped to: an API published to the portal.
resource "konnect_api" "my_api" {
  name        = "TF v3 portal team role API"
  description = "API used as the entity for a portal team role assignment."
  slug        = "tf-v3-portal-team-role-api"
  version     = "v1"
}

resource "konnect_api_publication" "my_publication" {
  api_id                     = konnect_api.my_api.id
  portal_id                  = konnect_portal.my_portal.id
  auto_approve_registrations = false
  visibility                 = "private"
}

resource "konnect_portal_team_role" "my_portal_team_role" {
  portal_id        = konnect_portal.my_portal.id
  team_id          = konnect_portal_team.my_portal_team.id
  role_name        = "API Viewer"
  entity_id        = konnect_api.my_api.id
  entity_type_name = "Services"
  entity_region    = "us"

  depends_on = [konnect_api_publication.my_publication]
}
