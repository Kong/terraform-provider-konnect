resource "konnect_gateway_route" "my_gatewayroute" {
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  destinations = [
    {
      ip   = "...my_ip..."
      port = 8
    }
  ]
  headers = {
    key = "value",
  }
  hosts = [
    "..."
  ]
  https_redirect_status_code = 307
  id                         = "...my_id..."
  methods = [
    "..."
  ]
  name          = "...my_name..."
  path_handling = "v0"
  paths = [
    "..."
  ]
  preserve_host = false
  protocols = [
    "tcp"
  ]
  regex_priority     = 9
  request_buffering  = true
  response_buffering = false
  service = {
    id = "...my_id..."
  }
  snis = [
    "..."
  ]
  sources = [
    {
      ip   = "...my_ip..."
      port = 0
    }
  ]
  strip_path = true
  tags = [
    "..."
  ]
}