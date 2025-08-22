resource "konnect_gateway_control_plane" "plugin_ratelimiting_advanced_cp" {
  name         = "Terraform Control Plane For RateLimitingAdvanced  Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_rate_limiting_advanced" "my_rate_limiting_advanced" {
  enabled = true

  config = {
    limit = [200]
    window_size = [1800]
    window_type = "fixed"
    namespace = "my-namespace"
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ratelimiting_advanced_cp.id
}
