resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Gateway Config Store"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_config_store" "my_configstore" {
  name = "demo-config-store"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_config_store_secret" "my_configstoresecret" {
  key   = "testing-key"
  value = "testing-value"

  config_store_id  = konnect_gateway_config_store.my_configstore.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

