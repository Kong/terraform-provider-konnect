resource "konnect_gateway_key_set" "my_keyset" {
  name             = "demo-keyset"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
