resource "konnect_mesh_control_plane" "my_meshcontrolplane" {
  description = "A control plane to handle traffic on development environment."
  labels = {
    key = "value",
  }
  name = "Test control plane"
}