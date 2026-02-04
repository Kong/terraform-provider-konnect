resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Demo API"
  description  = "My description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ace" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin"

  config = {
    match_policy = "required"
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}


resource "konnect_api" "api" {
  name = "my api"
  version = "1.0"
  slug    = "myapislug"
}

resource "konnect_api_implementation" "cp" {
  depends_on = [konnect_gateway_plugin_ace.my_plugin]
  api_id = konnect_api.api.id
  control_plane_reference = {
    control_plane = {
      id               = konnect_gateway_control_plane.my_konnect_cp.id
    }
  }
}