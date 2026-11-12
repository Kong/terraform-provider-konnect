resource "konnect_gateway_control_plane" "plugin_cors_cp" {
  name         = "Terraform Control Plane For CORS Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_cors" "my_cors" {
  enabled = true

  config = {
    origins = ["http://mockbin.com"]
    headers = ["http://mockbin.com"]
    methods = ["GET", "POST"]
  }

  control_plane_id = konnect_gateway_control_plane.plugin_cors_cp.id
}
