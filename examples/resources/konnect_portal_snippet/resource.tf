resource "konnect_portal_snippet" "my_portalsnippet" {
  content     = "# Welcome to My Snippet"
  description = "A custom page about developer portals"
  name        = "my-snippet"
  portal_id   = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  status      = "published"
  title       = "My Snippet"
  visibility  = "public"
}