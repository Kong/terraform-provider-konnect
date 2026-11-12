resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI MCP Oauth2 Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_mcp_oauth2" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin"

  config = {
    resource               = "https://example.com"
    authorization_servers  = ["https://auth.example.com"]
    introspection_endpoint = "https://introspection.example.com"
    client_id              = "client_id"
    client_secret          = "client_secret"
    claim_to_header = [
      {
        claim  = "sub"
        header = "X-User-Id"
    }]
    insecure_relaxed_audience_validation = true
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
