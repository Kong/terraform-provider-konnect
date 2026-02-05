resource "konnect_gateway_certificate" "my_gatewaycertificate" {
  cert             = "...my_cert..."
  cert_alt         = "...my_cert_alt..."
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 10
  id               = "...my_id..."
  key              = "...my_key..."
  key_alt          = "...my_key_alt..."
  tags = [
    "..."
  ]
  updated_at = 4
}