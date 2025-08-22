data "konnect_gateway_control_plane" "my_cp" {
  filter = {
    name = {
      eq = "Lookup Control Plane"
    }
  }
}