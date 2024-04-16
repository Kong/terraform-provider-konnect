resource "konnect_portal" "my_portal" {
  portal_id                 = data.konnect_portal_list.my_portallist.data[0].id
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "demo.example.com"
  is_public                 = false
  rbac_enabled              = false
}
