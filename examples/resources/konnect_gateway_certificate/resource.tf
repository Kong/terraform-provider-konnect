resource "konnect_gateway_certificate" "my_gatewaycertificate" {
  cert             = "...my_cert..."
  cert_alt         = "...my_cert_alt..."
  certificate_id   = "ddf3cdaa-3329-4961-822a-ce6dbd38eff7"
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  key              = "...my_key..."
  key_alt          = "...my_key_alt..."
}