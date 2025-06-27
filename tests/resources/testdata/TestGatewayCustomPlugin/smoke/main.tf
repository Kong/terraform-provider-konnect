resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Custom Plugin Test"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_custom_plugin" "custom_basic_auth" {
  name             = "basic-auth"
  instance_name    = "custom-plugin-test"
  config           = {}
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
