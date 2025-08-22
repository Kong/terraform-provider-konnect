resource "konnect_portal" "my_portal" {
  auto_approve_applications            = false
  auto_approve_developers              = false
  default_api_visibility               = "public"
  default_page_visibility              = "private"
  description                          = "My portal description"
  display_name                         = "My portal display name"

  name         = "My Portal"
  rbac_enabled = false
}