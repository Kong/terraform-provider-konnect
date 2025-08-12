resource "konnect_gateway_control_plane" "plugin_request_transformer_adv_cp" {
  name         = "Terraform Control Plane For RequestTransformerAdvanced Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_request_transformer_advanced" "my_request_transformer_advanced" {
  enabled = true

  config = {
    add = {
      headers = ["New-Header:Header Value"]
    }
    replace = {
      uri = "/status/(uri_captures[\"status\"])"
    }
  }

  control_plane_id = konnect_gateway_control_plane.plugin_request_transformer_adv_cp.id
}
