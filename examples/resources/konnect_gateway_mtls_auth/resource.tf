resource "konnect_gateway_mtls_auth" "my_gatewaymtlsauth" {
  ca_certificate = {
    id = "...my_id..."
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_id      = "f28acbfa-c866-4587-b688-0208ac24df21"
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  id               = "...my_id..."
  subject_name     = "...my_subject_name..."
  tags = [
    "..."
  ]
}