resource "konnect_api" "my_api" {
  description = "My API Description"
  name         = "MyAPI for Documents"
  slug         = "my-api-v1-publication"
  version      = "v1"
}

resource "konnect_portal" "my_portal_api_publication" {
  force_destroy                        = "true"
  name         = "My portal for api publication"
}

resource "konnect_api_publication" "my_apipublication" {
  api_id = konnect_api.my_api.id
  auto_approve_registrations = false
  portal_id                  = konnect_portal.my_portal_api_publication.id
  visibility                 = "private"
}
