resource "konnect_gateway_control_plane" "tfdemo" {
  name = "test-demo-tf"
  auth_type = "pki_client_certs"
  cloud_gateway = false
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

