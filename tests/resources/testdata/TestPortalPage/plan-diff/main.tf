resource "konnect_portal" "my_portal_for_page" {
  force_destroy                        = "true"
  name         = "My v3 portal name"
}

resource "konnect_portal_page" "my_portalpage" {
  content        = "# Welcome to My Page"
  description    = "A custom page about developer portals"
  portal_id      = konnect_portal.my_portal_for_page.id
  slug           = "/my-test-page"
  title          = "My Page"
}