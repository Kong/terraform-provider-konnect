resource "konnect_gateway_control_plane" "tfkic" {
  name         = "Terraform KIC Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER"
  auth_type    = "pinned_client_certs"

  proxy_urls = [
    {
      host     = "example.com",
      port     = 443,
      protocol = "https"
    }
  ]
}
