resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "My AI Gateway"
  name         = "my-ai-gateway"
}

resource "konnect_ai_gateway_mcp_server" "my_aigatewaymcpserver" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  upstream_server = {
    name         = "flights-mcp"
    display_name = "Flights MCP Server"
    enabled      = true
    config = {
      url = "https://mcp.internal.example.com"
      route = {
        paths = ["/flights-mcp"]
      }
      tools_cache_ttl_seconds = 10
    }
    tools = [
      {
        name        = "search-flights"
        description = "Search for available flights"
        host        = "https://api.example.com"
        scheme      = "https"
        method      = "GET"
        path        = "/flights"
        parameters = [
          {
            in   = "query"
            name = "origin"
          },
          {
            in   = "query"
            name = "destination"
          }
        ]
      }
    ]
  }
}
