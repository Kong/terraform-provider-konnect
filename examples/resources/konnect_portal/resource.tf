resource "konnect_portal" "my_portal" {
  auto_approve_applications            = false
  auto_approve_developers              = false
  custom_client_domain                 = "key-self-confidence.name"
  custom_domain                        = "measly-conservation.info"
  default_application_auth_strategy_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  description                          = "...my_description..."
  display_name                         = "...my_display_name..."
  force                                = "true"
  is_public                            = false
  labels = {
    key = "value",
  }
  name         = "...my_name..."
  rbac_enabled = true
}