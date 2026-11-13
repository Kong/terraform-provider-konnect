resource "konnect_gateway_control_plane" "plugin_ratelimiting_advanced_expressions_cp" {
  name         = "Terraform Control Plane For RateLimitingAdvanced Plugin With Expressions"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_rate_limiting_advanced" "my_rate_limiting_advanced_expressions" {
  enabled = true

  config = {
    limit       = [200]
    window_size = [1800]
    window_type = "fixed"
    namespace   = "my-namespace"
    header_name = "X-RateLimit-Limit"
    redis = {
      host = "redis.example.com"
      port = 6379
    }
  }

  # limit removed here - custom_key stays, to isolate whether just the leaf
  # expressions.limit is actually cleared server-side.
  expressions = {
    custom_key = "net.dst.ip"
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ratelimiting_advanced_expressions_cp.id
}
