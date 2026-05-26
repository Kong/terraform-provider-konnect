resource "konnect_gateway_control_plane" "plugin_ratelimiting_partial_cp" {
  name         = "Terraform Control Plane For RateLimiting Plugin With Partial"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_partial" "redis_ce_partial" {
  redis_ce = {
    name = "my_tf_redis_ce_partial"
    config = {
      host = "redis.example.com"
      port = 6379
    }
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ratelimiting_partial_cp.id
}

resource "konnect_gateway_plugin_rate_limiting" "my_rate_limiting" {
  enabled = true

  partials = [{
    id   = konnect_gateway_partial.redis_ce_partial.id
    path = "config.redis"
  }]

  config = {
    policy = "redis"
    hour   = 1000
    redis  = null
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ratelimiting_partial_cp.id
}
