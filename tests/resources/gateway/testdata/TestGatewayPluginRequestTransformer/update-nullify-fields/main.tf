resource "konnect_gateway_control_plane" "plugin_request_transformer_cp" {
  name         = "Terraform Control Plane For RequestTransformer Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_request_transformer" "my_request_transformer" {
  enabled = true

  config = {

    add = {
      headers = ["New-Header:Header Value"]
    }
    http_method = "GET"
  }

  control_plane_id = konnect_gateway_control_plane.plugin_request_transformer_cp.id
}