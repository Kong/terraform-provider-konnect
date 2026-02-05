resource "konnect_gateway_service" "my_gatewayservice" {
  ca_certificates = [
    "..."
  ]
  client_certificate = {
    id = "...my_id..."
  }
  connect_timeout  = 1826621898
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 10
  enabled          = true
  host             = "...my_host..."
  id               = "...my_id..."
  name             = "...my_name..."
  path             = "...my_path..."
  port             = 12962
  protocol         = "tls"
  read_timeout     = 985216672
  retries          = 9489
  tags = [
    "..."
  ]
  tls_sans = {
    dnsnames = [
      "..."
    ]
    uris = [
      "..."
    ]
  }
  tls_verify       = true
  tls_verify_depth = 49
  updated_at       = 8
  write_timeout    = 1780867234
}