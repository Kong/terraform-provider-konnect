resource "konnect_gateway_control_plane" "plugincp" {
  name         = "Terraform Control Plane For Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_rate_limiting" "my_nullable_rate_limiting" {
  enabled = true
  instance_name = "nullable-rl-plugin"
  config = {
    second = 5
    hour = 10000
    policy = "local"
    year = 1
  }
  // partials: to do - there is non-refresh plan diff which breaks tests.
  control_plane_id = konnect_gateway_control_plane.plugincp.id
}
