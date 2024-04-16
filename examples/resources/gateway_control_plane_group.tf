resource "konnect_gateway_control_plane" "tfcpg" {
  name         = "Terraform Control Plane Group"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE_GROUP"
  auth_type    = "pinned_client_certs"

  proxy_urls = []
}
