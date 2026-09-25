resource "konnect_gateway_control_plane" "plugin_ratelimiting_expressions_cp" {
  name         = "Terraform Control Plane For RateLimiting Plugin With Expressions"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_rate_limiting" "my_rate_limiting_expressions" {
  enabled = true

  config = {
    policy = "local"
    hour   = 1000
  }

  expressions = {
    custom_key = "net.dst.ip"
    day        = "300"
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ratelimiting_expressions_cp.id
}
