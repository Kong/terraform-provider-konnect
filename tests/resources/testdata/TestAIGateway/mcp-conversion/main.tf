resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - mcp-conversion"
  name         = "tf-test-aigw-mcp-conversion"
}

resource "konnect_ai_gateway_mcp_server" "my_aigatewaymcpserver_conversion" {
  provider = konnect-beta
  conversion_only = {
    config = {
      logging = {
        audits = true
      }
      max_request_body_size = 8388
      route = {
        paths = [
          "/foo"
        ]
        hosts = []
        protocols = [
          "https"
        ]
        https_redirect_status_code = 426
        regex_priority             = 0
        request_buffering          = true
        response_buffering         = true
        strip_path                 = true
      }
      url = "https://mcp.internal.kongair.com"
    }
    display_name = "TF Test MCP Conversion"
    enabled      = true
    name         = "tf-test-mcp-conversion"
    tools = [
      {
        description = "aaaaaaa"
        host        = "https://example.com"
        method      = "GET"
        name        = "search-examples"
        path        = "/examples"
        scheme      = "https"
      }
    ]
  }
  gateway_id = konnect_ai_gateway.my_aigateway.id
}
