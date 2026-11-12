resource "konnect_portal" "my_portal_for_snippet" {
  force_destroy = "true"
  name         = "My v3 portal name"
}

resource "konnect_portal_snippet" "my_portal_snippet" {
  content     = "# Welcome to My Snippet"
  description = "A custom snippet about developer portals"
  name        = "my-snippet"
  title       = "My Snippet"
  portal_id   = konnect_portal.my_portal_for_snippet.id
}
