resource "konnect_gateway_service" "my_gatewayservice" {
  ca_certificates = [
    "..."
  ]
  client_certificate = {
    id = "...my_id..."
  }
  connect_timeout  = 9
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = true
  host             = "...my_host..."
  id               = "...my_id..."
  name             = "...my_name..."
  path             = "...my_path..."
  port             = 2
  protocol         = "tls"
  read_timeout     = 5
  retries          = 3
  tags = [
    "..."
  ]
  tls_verify       = true
  tls_verify_depth = 8
  write_timeout    = 9
}