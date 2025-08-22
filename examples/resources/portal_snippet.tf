resource "konnect_portal_snippet" "my_portalsnippet" {
  content     = "# Welcome to My Snippet"
  description = "A custom page about developer portals"
  name        = "my-snippet"
  portal_id   = konnect_portal.my_portal.id
  status      = "published"
  title       = "My Snippet"
  visibility  = "public"
}