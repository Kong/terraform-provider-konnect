resource "konnect_portal" "my_portal" {
  authentication_enabled               = false
  auto_approve_applications            = false
  auto_approve_developers              = false
  default_api_visibility               = "public"
  default_application_auth_strategy_id = "e7d77a5f-c5f5-49db-9b2f-cabb4401add8"
  default_page_visibility              = "private"
  description                          = "...my_description..."
  display_name                         = "...my_display_name..."
  force_destroy                        = "false"
  labels = {
    key = "value"
  }
  name         = "...my_name..."
  rbac_enabled = true
}