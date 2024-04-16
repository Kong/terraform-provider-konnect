resource "konnect_portal_auth" "my_portalauth" {
  portal_id               = data.konnect_portal_list.my_portallist.data[0].id
  basic_auth_enabled      = true
  konnect_mapping_enabled = true

  oidc_auth_enabled         = false
  oidc_team_mapping_enabled = false
}
