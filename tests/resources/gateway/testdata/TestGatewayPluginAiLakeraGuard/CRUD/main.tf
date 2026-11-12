resource "konnect_gateway_control_plane" "plugin_ai_lakera_guard_cp" {
  name         = "Terraform Control Plane For AI Lakera Guard Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_lakera_guard" "my_ai_lakera_guard" {
  enabled = true

  config = {
    api_key        = "test-api-key-12345"
    guarding_mode  = "INPUT"
    text_source    = "last_message"
    timeout        = 5000
  }

  control_plane_id = konnect_gateway_control_plane.plugin_ai_lakera_guard_cp.id
}
