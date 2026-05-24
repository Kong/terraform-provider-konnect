resource "konnect_gateway_control_plane" "plugin_ratelimiting_advanced_partial_cp" {
  name         = "Terraform Control Plane For RateLimitingAdvanced Plugin With Partial"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_partial" "redis_ee_partial" {
  redis_ee = {
    name = "my_tf_redis_ee_partial"
    config = {
      host = "redis.example.com"
      port = 6379
    }
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ratelimiting_advanced_partial_cp.id
}

resource "konnect_gateway_plugin_rate_limiting_advanced" "my_rate_limiting_advanced" {
  enabled = true

  partials = [{
    id   = konnect_gateway_partial.redis_ee_partial.id
    path = "config.redis"
  }]

  config = {
    limit       = [200]
    window_size = [1800]
    window_type = "fixed"
    namespace   = "my-namespace"
  }

  lifecycle {
    ignore_changes = [ "config.redis.cluster_max_redirections" ]
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ratelimiting_advanced_partial_cp.id
}
