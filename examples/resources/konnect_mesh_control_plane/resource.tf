resource "konnect_mesh_control_plane" "my_meshcontrolplane" {
  description = "A control plane to handle traffic on development environment."
  features = [
    {
      hostname_generator_creation = {
        enabled = false
      }
      mesh_creation = {
        enabled = false
      }
      type = "MeshCreation"
    }
  ]
  labels = {
    key = "value"
  }
  name = "Test control plane"
}