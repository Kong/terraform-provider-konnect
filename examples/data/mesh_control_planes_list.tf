data "konnect_mesh_control_planes" "my_meshcontrolplanes" {
  filter = {
    name = {
      contains = "my-mesh"
    }
    labels = {
      eq = "team:platform"
    }
  }
}