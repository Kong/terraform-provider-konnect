resource "konnect_gateway_route_expression" "my_gatewayrouteexpression" {
  control_plane_id           = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  expression                 = "...my_expression..."
  https_redirect_status_code = 307
  id                         = "...my_id..."
  name                       = "...my_name..."
  path_handling              = "v0"
  preserve_host              = true
  priority                   = 6
  protocols = [
    "http"
  ]
  request_buffering  = false
  response_buffering = true
  service = {
    id = "...my_id..."
  }
  strip_path = true
  tags = [
    "..."
  ]
}