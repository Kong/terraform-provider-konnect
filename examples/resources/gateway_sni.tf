resource "konnect_gateway_sni" "my_sni" {
  name = "example.com"
  certificate = {
    id = konnect_gateway_certificate.my_certificate.id
  }
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
