resource "konnect_gateway_plugin_header_cert_auth" "my_gatewaypluginheadercertauth" {
  config = {
    allow_partial_chain    = false
    anonymous              = "...my_anonymous..."
    authenticated_group_by = "DN"
    ca_certificates = [
      "..."
    ]
    cache_ttl                 = 5.31
    cert_cache_ttl            = 6.32
    certificate_header_format = "url_encoded"
    certificate_header_name   = "...my_certificate_header_name..."
    consumer_by = [
      "username"
    ]
    default_consumer      = "...my_default_consumer..."
    http_proxy_host       = "...my_http_proxy_host..."
    http_proxy_port       = 1064
    http_timeout          = 5.78
    https_proxy_host      = "...my_https_proxy_host..."
    https_proxy_port      = 10840
    revocation_check_mode = "SKIP"
    secure_source         = true
    skip_consumer_lookup  = true
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
    "https"
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