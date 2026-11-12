resource "konnect_gateway_control_plane" "pluginopentelemetrycp" {
  name         = "Terraform Control Plane For OpenTelemetry Plugin Access Logs"
  description  = "Test control plane for access_logs_endpoint drift fix"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_opentelemetry" "my_opentelemetry" {
  enabled = true
  config = {
    traces_endpoint = "http://localhost:4318/v1/traces"

    # Using access_logs.endpoint (not access_logs_endpoint)
    # The API will auto-populate access_logs_endpoint from this value
    # This test verifies there's no drift on subsequent plans
    access_logs = {
      endpoint = "http://localhost:4318/v1/logs"
    }
  }
  control_plane_id = konnect_gateway_control_plane.pluginopentelemetrycp.id
}

