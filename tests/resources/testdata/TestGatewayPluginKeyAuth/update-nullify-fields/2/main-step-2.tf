resource "konnect_gateway_control_plane" "plugin_key_auth_cp" {
  name         = "Terraform Control Plane For Key Auth Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_key_auth" "my_key_auth" {
  enabled = true

  config = {
    key_names = ["apikey"]
  }

  control_plane_id = konnect_gateway_control_plane.plugin_key_auth_cp.id
}