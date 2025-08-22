resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Lookup Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

data "konnect_gateway_control_plane" "my_cp" {
  filter = {
    name = {
      eq = "Lookup Control Plane"
    }
  }
  depends_on = [konnect_gateway_control_plane.tfdemo]
}

output "control_plane" {
  value = data.konnect_gateway_control_plane.my_cp.name
}
