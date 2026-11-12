resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For AI MCP Proxy Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_ai_mcp_proxy" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin"

  config = {
    mode = "conversion-listener"
    tools = [
      {
        description = "Get users"
        method      = "GET"
        path        = "/marketplace/users"

        annotations = {
          title = "Get users"
        }
        parameters = [
          {
            name     = "id"
            in       = "query"
            required = false

            schema = {
              type = "string"
            }
            description = "Optional user ID"
        }]
      },
      {
        description = "Get orders for a user - require parameter"
        method      = "GET"
        path        = "/marketplace/orders"

        annotations = {
          title = "Get users orders"
        }
        parameters = [
          {
            name     = "userid"
            in       = "query"
            required = true

            schema = {
              type = "string"
            }
            description = "User ID to filter orders"
        }]
      },

      {
        description = "Get orders for a user - body"
        method      = "GET"
        path        = "/marketplace/orders"

        annotations = {
          title = "Get users orders"
        }
        request_body = jsonencode({
          "content": {"application/json":{"schema":{"type":"object","properties":{"userid":{"type":"string"}},"required":["userid"]}}}
        })
    }]

    server = {
      timeout = 60000
    }
    max_request_body_size = 8192
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
