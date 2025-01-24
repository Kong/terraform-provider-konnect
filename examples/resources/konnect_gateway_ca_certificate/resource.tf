resource "konnect_gateway_ca_certificate" "my_gatewaycacertificate" {
  cert             = "...my_cert..."
  cert_digest      = "...my_cert_digest..."
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  id               = "...my_id..."
  tags = [
    "..."
  ]
}