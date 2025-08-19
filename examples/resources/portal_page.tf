resource "konnect_portal_page" "my_portalpage" {
  content        = "# Welcome to My Page"
  description    = "A custom page about developer portals"
  portal_id      = konnect_portal.my_portal.id
  slug           = "/my-page"
  status         = "published"
  title          = "My Page"
  visibility     = "public"
}