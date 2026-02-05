resource "konnect_gateway_plugin_tls_metadata_headers" "my_gatewayplugintlsmetadataheaders" {
  config = {
    client_cert_fingerprint_header_name = "...my_client_cert_fingerprint_header_name..."
    client_cert_header_name             = "...my_client_cert_header_name..."
    client_cert_issuer_dn_header_name   = "...my_client_cert_issuer_dn_header_name..."
    client_cert_subject_dn_header_name  = "...my_client_cert_subject_dn_header_name..."
    client_serial_header_name           = "...my_client_serial_header_name..."
    inject_client_cert_details          = false
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 6
  enabled          = true
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
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
  protocols = [
    "grpcs"
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
  updated_at = 2
}