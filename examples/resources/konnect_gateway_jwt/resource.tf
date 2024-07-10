resource "konnect_gateway_jwt" "my_gatewayjwt" {
  algorithm        = "ES512"
  consumer_id      = "f28acbfa-c866-4587-b688-0208ac24df21"
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  jwt_id           = "4a7f5faa-8c96-46d6-8214-c87573ef2ac4"
  key              = "...my_key..."
  rsa_public_key   = "...my_rsa_public_key..."
  secret           = "...my_secret..."
}