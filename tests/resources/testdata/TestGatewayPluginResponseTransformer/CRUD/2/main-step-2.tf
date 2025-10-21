resource "konnect_gateway_control_plane" "plugincp" {
  name         = "Terraform Control Plane For Response transformer Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_response_transformer" "my_response_transformer" {
  enabled = true

  config = {
    remove = {
      headers = ["New-Header:Header Value"]
    }
    add = {
      headers = []
    }
  }

  control_plane_id = konnect_gateway_control_plane.plugincp.id
}
