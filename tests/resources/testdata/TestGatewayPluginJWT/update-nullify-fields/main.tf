resource "konnect_gateway_control_plane" "plugin_jwt_cp" {
  name         = "Terraform Control Plane For JWT Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_jwt" "my_jwt" {
  enabled = true

  config = {
    cookie_names = ["jwt"]
    anonymous = "anon-user"
  }

  control_plane_id = konnect_gateway_control_plane.plugin_jwt_cp.id
}
