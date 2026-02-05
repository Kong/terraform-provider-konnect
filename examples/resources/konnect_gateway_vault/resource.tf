resource "konnect_gateway_vault" "my_gatewayvault" {
  config           = "{ \"see\": \"documentation\" }"
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 4
  description      = "...my_description..."
  id               = "...my_id..."
  name             = "...my_name..."
  prefix           = "...my_prefix..."
  tags = [
    "..."
  ]
  updated_at = 4
}