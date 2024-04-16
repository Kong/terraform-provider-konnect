resource "konnect_gateway_target" "my_target" {
  target           = "one.example.com:9000"
  weight           = 1
  upstream_id      = konnect_gateway_upstream.my_upstream.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
