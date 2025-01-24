resource "konnect_gateway_plugin_ldap_auth_advanced" "my_gatewaypluginldapauthadvanced" {
  config = {
    anonymous = "...my_anonymous..."
    attribute = "...my_attribute..."
    base_dn   = "...my_base_dn..."
    bind_dn   = "...my_bind_dn..."
    cache_ttl = 2.41
    consumer_by = [
      "custom_id"
    ]
    consumer_optional      = true
    group_base_dn          = "...my_group_base_dn..."
    group_member_attribute = "...my_group_member_attribute..."
    group_name_attribute   = "...my_group_name_attribute..."
    groups_required = [
      "..."
    ]
    header_type        = "...my_header_type..."
    hide_credentials   = true
    keepalive          = 3.29
    ldap_host          = "...my_ldap_host..."
    ldap_password      = "...my_ldap_password..."
    ldap_port          = 3.46
    ldaps              = false
    log_search_results = false
    realm              = "...my_realm..."
    start_tls          = true
    timeout            = 8.28
    verify_ldap_host   = true
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
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