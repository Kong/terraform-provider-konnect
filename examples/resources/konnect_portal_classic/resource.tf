resource "konnect_portal_classic" "my_portalclassic" {
  auto_approve_applications            = false
  auto_approve_developers              = false
  custom_client_domain                 = "frugal-juggernaut.name"
  custom_domain                        = "smoggy-draw.info"
  default_application_auth_strategy_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  description                          = "...my_description..."
  display_name                         = "...my_display_name..."
  force                                = "false"
  is_public                            = true
  labels = {
    key = "value"
  }
  name         = "...my_name..."
  rbac_enabled = true
}