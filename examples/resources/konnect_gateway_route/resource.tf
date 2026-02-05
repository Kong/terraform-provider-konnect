resource "konnect_gateway_route" "my_gatewayroute" {
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 8
  destinations = [
    {
      ip   = "...my_ip..."
      port = 47938
    }
  ]
  headers = {
    key = [
      # ...
    ]
  }
  hosts = [
    "..."
  ]
  https_redirect_status_code = 308
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
      port = 633
    }
  ]
  strip_path = true
  tags = [
    "..."
  ]
  updated_at = 6
}