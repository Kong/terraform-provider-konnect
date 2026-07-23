resource "konnect_mesh_control_plane" "tfdemo" {
  name        = "Lookup Mesh Control Plane"
  description = "This is a sample description"
}

data "konnect_mesh_control_planes" "my_meshcontrolplanes" {
  filter = {
    name = {
      eq = "Lookup Mesh Control Plane"
    }
  }
  depends_on = [konnect_mesh_control_plane.tfdemo]
}
