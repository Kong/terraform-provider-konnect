resource "konnect_gateway_target" "my_gatewaytarget" {
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  id               = "...my_id..."
  tags = [
    "..."
  ]
  target = "...my_target..."
  upstream = {
    id = "...my_id..."
  }
  upstream_id = "5a078780-5d4c-4aae-984a-bdc6f52113d8"
  weight      = 5
}