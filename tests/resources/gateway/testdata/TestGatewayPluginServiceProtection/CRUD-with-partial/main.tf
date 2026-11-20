resource "konnect_gateway_control_plane" "plugin_svc_protection_cp" {
  name         = "Terraform Control Plane For ServiceProtection Plugin With Partial"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_partial" "redis_ee_partial" {
  redis_ee = {
    name = "my_tf_svcprotect_redis_ee_partial"
    config = {
      host = "redis.example.com"
      port = 6379
    }
  }

  control_plane_id = konnect_gateway_control_plane.plugin_svc_protection_cp.id
}

resource "konnect_gateway_plugin_service_protection" "my_service_protection" {
  enabled       = true
  instance_name = "example-service-protection"
  protocols     = ["http"]

  partials = [{
    id   = konnect_gateway_partial.redis_ee_partial.id
    path = "config.redis"
  }]

  config = {
    disable_penalty     = false
    error_code          = 503
    error_message       = "Service temporarily unavailable"
    hide_client_headers = true
    namespace           = "example"
    limit               = [10]
    strategy            = "redis"
    sync_rate           = 0
    window_size         = [60]
    redis = {}
  }

  control_plane_id = konnect_gateway_control_plane.plugin_svc_protection_cp.id
}
