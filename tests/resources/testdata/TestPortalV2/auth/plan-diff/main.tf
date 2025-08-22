resource "konnect_portal_classic" "my_portal" {
  name                      = "My v2 Developer Portal for auth"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "tf-example-domain-portal-auth.example.com"
  is_public                 = false
  rbac_enabled              = false
}

resource "konnect_portal_auth" "my_portalauth" {
  portal_id               = konnect_portal_classic.my_portal.id
  basic_auth_enabled      = true
  konnect_mapping_enabled = true

  oidc_auth_enabled         = false
  oidc_team_mapping_enabled = false
}
