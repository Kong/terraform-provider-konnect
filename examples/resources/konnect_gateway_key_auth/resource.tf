resource "konnect_gateway_key_auth" "my_gatewaykeyauth" {
  consumer = {
    id = "...my_id..."
  }
  consumer_id      = "f28acbfa-c866-4587-b688-0208ac24df21"
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  id               = "...my_id..."
  key              = "...my_key..."
  tags = [
    "..."
  ]
}