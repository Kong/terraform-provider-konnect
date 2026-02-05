resource "konnect_portal_page" "my_portalpage" {
  content        = "# Welcome to My Page"
  description    = "A custom page about developer portals"
  parent_page_id = "824a28d9-7024-426a-aed4-03b504521824"
  portal_id      = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  slug           = "/my-page"
  status         = "published"
  title          = "My Page"
  visibility     = "public"
}