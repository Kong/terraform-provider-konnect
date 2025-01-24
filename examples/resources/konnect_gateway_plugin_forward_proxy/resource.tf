resource "konnect_gateway_plugin_forward_proxy" "my_gatewaypluginforwardproxy" {
  config = {
    auth_password    = "...my_auth_password..."
    auth_username    = "...my_auth_username..."
    http_proxy_host  = "...my_http_proxy_host..."
    http_proxy_port  = 61130
    https_proxy_host = "...my_https_proxy_host..."
    https_proxy_port = 38011
    https_verify     = false
    proxy_scheme     = "http"
    x_headers        = "transparent"
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = false
  id               = "...my_id..."
  instance_name    = "...my_instance_name..."
  ordering = {
    after = {
      access = [
        "..."
      ]
    }
    before = {
      access = [
        "..."
      ]
    }
  }
  protocols = [
    "ws"
  ]
  route = {
    id = "...my_id..."
  }
  service = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
}