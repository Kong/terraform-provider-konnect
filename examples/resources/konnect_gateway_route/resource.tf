resource "konnect_gateway_route" "my_gatewayroute" {
  control_plane_id           = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  https_redirect_status_code = "301"
  name                       = "Jo Steuber"
  path_handling              = "v1"
  preserve_host              = true
  regex_priority             = 7
  request_buffering          = false
  response_buffering         = false
  route_id                   = "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7"
  strip_path                 = false
}