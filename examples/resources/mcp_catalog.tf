#example of creating a MCP in the catalog with a version and an implementation that uses the AI Gateway and MCP Server
resource "konnect_catalog_mcp" "test_mcp" {
  name         = "test-mcp-tf"
  display_name = "Test MCP"
  description  = "MCP for testing"
  labels = {
    env = "test"
  }
}

#example of creating a version for the MCP in the catalog with a resource
resource "konnect_catalog_mcp_version" "test_mcp_version" {
  mcp_id  = konnect_catalog_mcp.test_mcp.id
  version = "1.0.0"
  resources = [
    {
      name        = "config-file"
      uri         = "file:///etc/mcp/config.json"
      title       = "MCP Configuration"
      description = "Configuration file for the MCP"
      mime_type   = "application/json"
      size        = 1234
    }
  ]
}

#example of creating an AI Gateway  to be used in the MCP implementation(currently in beta-provider as of 02/09/2026)
resource "konnect_ai_gateway" "test_ai_gateway_mcp" {
  provider     = konnect-beta
  name         = "test-ai-gateway-mcp-impl"
  display_name = "Test AI Gateway for MCP"
}

#example of creating an MCP Server to be used in the MCP implementation(currently in beta-provider as of 02/09/2026)
resource "konnect_ai_gateway_mcp_server" "test_mcp_server" {
  provider   = konnect-beta
  gateway_id = konnect_ai_gateway.test_ai_gateway_mcp.id
  upstream_server = {
    name         = "test-mcp-server"
    display_name = "Test MCP Server"
    config = {
      url = "https://mcp.example.com"
      tools_cache_ttl_seconds = 300
      route = {
        paths = ["/mcp"]
        hosts = []
      }
    }
  }
}

#example of creating a MCP implementation in the catalog that uses the AI Gateway and MCP Server
resource "konnect_catalog_mcp_implementation" "test_mcp_impl" {
  mcp_id = konnect_catalog_mcp.test_mcp.id
  create_catalog_mcp_gateway_implementation = {
    implementation = {
      config = {
        gateway_control_plane_id = konnect_ai_gateway.test_ai_gateway_mcp.id
        gateway_mcp_server_id    = konnect_ai_gateway_mcp_server.test_mcp_server.id
      }
      type = "ai-gateway"
    }
  }
}