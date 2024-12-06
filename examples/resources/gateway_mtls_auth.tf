resource "konnect_gateway_mtls_auth" "my_mtlsauth" {
  subject_name     = "example.com"
  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
