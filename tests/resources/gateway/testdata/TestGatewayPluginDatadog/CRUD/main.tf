resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane - Plugin Datadog"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_datadog" "my_datadog" {
  enabled = true
  config = {
    host = "127.0.0.1"
    port = 8125
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
