resource "konnect_portal" "my_portal" {
  auto_approve_applications            = false
  auto_approve_developers              = false
  custom_client_domain                 = "haunting-lieu.info"
  custom_domain                        = "square-silk.biz"
  default_application_auth_strategy_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  description                          = "...my_description..."
  display_name                         = "...my_display_name..."
  force                                = "true"
  is_public                            = true
  name                                 = "Hope Hodkiewicz"
  portal_id                            = "e9a891ba-5c14-4e32-81f1-6f4324b113e9"
  rbac_enabled                         = false
}