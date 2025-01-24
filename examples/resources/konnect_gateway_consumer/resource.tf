resource "konnect_gateway_consumer" "my_gatewayconsumer" {
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  custom_id        = "...my_custom_id..."
  id               = "...my_id..."
  tags = [
    "..."
  ]
  username = "...my_username..."
}