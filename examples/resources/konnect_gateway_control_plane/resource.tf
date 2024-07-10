resource "konnect_gateway_control_plane" "my_gatewaycontrolplane" {
  auth_type     = "pinned_client_certs"
  cloud_gateway = false
  cluster_type  = "CLUSTER_TYPE_CONTROL_PLANE"
  description   = "A test control plane for exploration."
  name          = "Test Control Plane"
  proxy_urls = [
    {
      host     = "...my_host..."
      port     = 10
      protocol = "...my_protocol..."
    },
  ]
}