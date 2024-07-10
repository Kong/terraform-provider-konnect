resource "konnect_gateway_service" "my_gatewayservice" {
  connect_timeout  = 1
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = false
  host             = "...my_host..."
  name             = "Sabrina Buckridge"
  path             = "...my_path..."
  port             = 1
  protocol         = "tls_passthrough"
  read_timeout     = 2
  retries          = 9
  service_id       = "7fca84d6-7d37-4a74-a7b0-93e576089a41"
  tls_verify       = true
  tls_verify_depth = 6
  write_timeout    = 5
}