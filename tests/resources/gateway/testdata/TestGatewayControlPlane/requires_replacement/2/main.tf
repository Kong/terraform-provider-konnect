resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane"
  description  = "This is a sample description"

  # Changing the cluster type requires a replacement
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE_GROUP"
}
