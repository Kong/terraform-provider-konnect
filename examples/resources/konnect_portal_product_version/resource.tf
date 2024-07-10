resource "konnect_portal_product_version" "my_portalproductversion" {
  application_registration_enabled = false
  auth_strategy_ids = [
    "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
  ]
  auto_approve_registration = false
  deprecated                = false
  notify_developers         = false
  portal_id                 = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  product_version_id        = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  publish_status            = "published"
}