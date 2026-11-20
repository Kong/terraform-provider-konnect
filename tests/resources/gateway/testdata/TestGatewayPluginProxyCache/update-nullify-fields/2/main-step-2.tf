resource "konnect_gateway_control_plane" "plugin_proxy_cache_cp" {
  name         = "Terraform Control Plane For ProxyCache Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_proxy_cache" "my_proxy_cache" {
  enabled = true

  config = {
    response_code = [200]
    request_method = ["GET"]
    content_type = ["text/plain", "application/json"]
    cache_ttl = 300
    strategy = "memory"
  }

  control_plane_id = konnect_gateway_control_plane.plugin_proxy_cache_cp.id
}
