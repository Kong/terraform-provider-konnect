resource "konnect_gateway_plugin_mtls_auth" "my_gatewaypluginmtlsauth" {
  config = {
    allow_partial_chain    = false
    anonymous              = "...my_anonymous..."
    authenticated_group_by = "DN"
    ca_certificates = [
      "..."
    ]
    cache_ttl      = 5.53
    cert_cache_ttl = 4.84
    consumer_by = [
      "custom_id"
    ]
    default_consumer      = "...my_default_consumer..."
    http_proxy_host       = "...my_http_proxy_host..."
    http_proxy_port       = 30482
    http_timeout          = 4.02
    https_proxy_host      = "...my_https_proxy_host..."
    https_proxy_port      = 17238
    revocation_check_mode = "STRICT"
    send_ca_dn            = true
    skip_consumer_lookup  = true
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
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
    "tls_passthrough"
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