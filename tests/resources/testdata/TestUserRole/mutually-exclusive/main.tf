resource "konnect_user_role" "my_userrole" {
  entity_id        = "5f2d8a1e-0f6b-4b3e-9a4d-1c2b3a4d5e6f"
  entity_ids       = ["7c9e4f2a-3b1d-4c5e-8f6a-9b0c1d2e3f4a"]
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Viewer"
  user_id          = "df5404e1-5c36-414b-a40b-903387226475"
}
