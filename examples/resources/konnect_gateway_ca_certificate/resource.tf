resource "konnect_gateway_ca_certificate" "my_gatewaycacertificate" {
  ca_certificate_id = "3c31f18a-f27a-4f9b-8cd4-bf841554612f"
  cert              = "...my_cert..."
  cert_digest       = "...my_cert_digest..."
  control_plane_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
}