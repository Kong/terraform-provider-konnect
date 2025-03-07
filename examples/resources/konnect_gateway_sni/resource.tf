resource "konnect_gateway_sni" "my_gatewaysni" {
  certificate = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  id               = "...my_id..."
  name             = "...my_name..."
  tags = [
    "..."
  ]
}