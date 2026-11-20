resource "konnect_gateway_control_plane" "plugin_canary_cp" {
  name         = "Terraform Control Plane For Canary Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_canary" "my_canary" {
  enabled = true

  config = {
    duration = 3600
    start = 0
    steps = 5  
    upstream_host = "my_host"
    upstream_port = 80
    upstream_uri  = "test"
    hash = "allow"
  }

  control_plane_id = konnect_gateway_control_plane.plugin_canary_cp.id
}
