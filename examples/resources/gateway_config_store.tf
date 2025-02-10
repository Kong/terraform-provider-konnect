resource "konnect_gateway_config_store" "my_configstore" {
  name  = "tf-config-store"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_vault" "my_vault" {
  name   = "konnect"
  prefix = "my-konnect-vault"
  config = jsonencode({
    config_store_id = konnect_gateway_config_store.my_configstore.id
  })
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
