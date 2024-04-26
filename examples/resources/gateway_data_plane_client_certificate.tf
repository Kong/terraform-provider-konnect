resource "konnect_gateway_data_plane_client_certificate" "my_cert" {
  cert             = file("./tls.crt")
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
