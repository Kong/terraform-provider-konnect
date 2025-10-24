resource "konnect_gateway_control_plane" "plugin_ace_cp" {
  name         = "Terraform Control Plane For ACE Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ace" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin"

  config = {
    match_policy = "required"
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ace_cp.id
}
