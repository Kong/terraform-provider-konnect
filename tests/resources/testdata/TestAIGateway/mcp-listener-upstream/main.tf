resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - mcp-listener"
  name         = "tf-test-aigw-mcp-listener"
}

resource "konnect_ai_gateway_mcp_server" "my_aigatewaymcpserver_upstream" {
  provider = konnect-beta
  upstream_server = {
    config = {
      tools_cache_ttl_seconds = 10
      logging = {
        audits   = true
      }
      max_request_body_size = 8388
      route = {
        paths = [
          "/foo"
        ]
        hosts = []
        methods = []
        protocols = [
          "https"
        ]
        tags = []
        https_redirect_status_code = 426
        regex_priority     = 0
        request_buffering  = true
        response_buffering = true
        strip_path         = true
      }
      url = "https://mcp.internal.kongair.com"
    }
    display_name = "TF Test MCP Conversion"
    enabled      = true
    name = "tf-test-mcp-upstream"
    tools = [
      {
        description = "aaaaaaa"
        host        = "https://example.com"
        method      = "GET"
        name        = "search-examples"
        path        = "/examples"
        scheme      = "https"
        annotations = {
          destructive_hint = true
        }
        parameters = [
          {
            in          = "query"
            name        = "origin"
          }
        ]
      }
    ]
  }
  gateway_id = konnect_ai_gateway.my_aigateway.id
}

resource "konnect_ai_gateway_mcp_server" "my_aigatewaymcpserver_listener" {
  provider = konnect-beta
  listener = {
    access = {
      consumer = {
        acls = {
          allow = [
            "my-test-allow-consumer"
          ]
        }
      }
    }
    config = {
      logging = {
        audits   = false
        payloads = false
      }
      max_request_body_size = 8388
      route = {
        hosts = [
          "foo.example.com"
        ]
        https_redirect_status_code = 426
        regex_priority             = 0
        request_buffering          = true
        response_buffering         = true
        strip_path                 = true
      }
    }
    display_name = "TF Test MCP Listener"
    enabled      = true
    name         = "tf-test-mcp-listener"
    sources = [
      konnect_ai_gateway_mcp_server.my_aigatewaymcpserver_upstream.upstream_server.name,
    ]
  }
  gateway_id = konnect_ai_gateway.my_aigateway.id
}
