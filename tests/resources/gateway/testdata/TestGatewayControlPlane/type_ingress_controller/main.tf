resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "KIC Terraform Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER"
}
