resource "konnect_gateway_control_plane" "plugin_basicauth_cp" {
  name         = "Terraform Control Plane For BasicAuth Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_basic_auth" "my_basic_auth" {
  enabled = true

  config = {
  }

  control_plane_id = konnect_gateway_control_plane.plugin_basicauth_cp.id
}
