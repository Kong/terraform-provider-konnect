resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Gateway Plugin Response Transformer adv Test CP"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_response_transformer_advanced" "my_response_transformer" {
  enabled = true

  config = {
    dots_in_keys = false
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}