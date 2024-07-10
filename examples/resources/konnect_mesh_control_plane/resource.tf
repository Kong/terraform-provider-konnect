resource "konnect_mesh_control_plane" "my_meshcontrolplane" {
  cp_id       = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  description = "A control plane to handle traffic on development environment."
  name        = "Test control plane"
}