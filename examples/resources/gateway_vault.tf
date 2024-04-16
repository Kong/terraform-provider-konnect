resource "konnect_gateway_vault" "my_vault" {
  name   = "env"
  prefix = "my-env-vault"
  config = jsonencode({
    prefix = "bbbb"
  })
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
