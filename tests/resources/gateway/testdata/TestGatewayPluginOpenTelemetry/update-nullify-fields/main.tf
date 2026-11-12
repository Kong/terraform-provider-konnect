resource "konnect_gateway_control_plane" "pluginopentelemetrycp" {
  name         = "Terraform Control Plane For OpenTelemetry Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}


resource "konnect_gateway_plugin_opentelemetry" "my_opentelemetry" {
  enabled = true
  config = {
    headers = {
      X-Auth-Token = jsonencode("secret-token")
    }
    traces_endpoint = "http://localhost:4318/v1/traces"
    logs_endpoint = "http://localhost:4318/v1/logs"
    endpoint = "http://localhost:4318/v1/traces"
  }
  control_plane_id = konnect_gateway_control_plane.pluginopentelemetrycp.id
}
