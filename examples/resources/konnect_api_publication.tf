resource "konnect_api_publication" "my_apipublication" {
  api_id                     = konnect_api.my_api.id
  auto_approve_registrations = false
  portal_id                  = konnect_portal.my_portal.id
  visibility                 = "private"
}
