resource "konnect_gateway_plugin_ldap_auth" "my_gatewaypluginldapauth" {
  config = {
    anonymous        = "...my_anonymous..."
    attribute        = "...my_attribute..."
    base_dn          = "...my_base_dn..."
    cache_ttl        = 0.75
    header_type      = "...my_header_type..."
    hide_credentials = false
    keepalive        = 2.29
    ldap_host        = "...my_ldap_host..."
    ldap_port        = 57764
    ldaps            = true
    realm            = "...my_realm..."
    start_tls        = false
    timeout          = 2.42
    verify_ldap_host = true
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
    "wss"
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