resource "konnect_team_role" "my_teamrole" {
  entity_id        = "5f2d8a1e-0f6b-4b3e-9a4d-1c2b3a4d5e6f"
  entity_ids       = ["7c9e4f2a-3b1d-4c5e-8f6a-9b0c1d2e3f4a"]
  entity_region    = "us"
  entity_type_name = "Control Planes"
  role_name        = "Viewer"
  team_id          = "a1b2c3d4-0000-4000-8000-000000000001"
}
