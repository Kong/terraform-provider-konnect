resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "File Log Stdout Test CP"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_file_log" "test" {
  enabled = true

  config = {
    path = "/dev/stdout"
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
