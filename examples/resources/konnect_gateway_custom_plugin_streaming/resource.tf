resource "konnect_gateway_custom_plugin_streaming" "my_gatewaycustompluginstreaming" {
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 10
  handler          = "...my_handler..."
  id               = "...my_id..."
  name             = "...my_name..."
  schema           = "...my_schema..."
  tags = [
    "..."
  ]
  updated_at = 1
}