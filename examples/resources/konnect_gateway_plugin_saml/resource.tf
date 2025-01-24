resource "konnect_gateway_plugin_saml" "my_gatewaypluginsaml" {
  config = {
    anonymous               = "...my_anonymous..."
    assertion_consumer_path = "...my_assertion_consumer_path..."
    idp_certificate         = "...my_idp_certificate..."
    idp_sso_url             = "...my_idp_sso_url..."
    issuer                  = "...my_issuer..."
    nameid_format           = "Unspecified"
    redis = {
      cluster_max_redirections = 0
      cluster_nodes = [
        {
          ip   = "...my_ip..."
          port = 21415
        }
      ]
      connect_timeout       = 1914874679
      connection_is_proxied = true
      database              = 7
      host                  = "...my_host..."
      keepalive_backlog     = 2023529059
      keepalive_pool_size   = 1633101853
      password              = "...my_password..."
      port                  = 6907
      prefix                = "...my_prefix..."
      read_timeout          = 1468960257
      send_timeout          = 1619402496
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 44971
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "master"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      socket            = "...my_socket..."
      ssl               = true
      ssl_verify        = false
      username          = "...my_username..."
    }
    request_digest_algorithm          = "SHA1"
    request_signature_algorithm       = "SHA384"
    request_signing_certificate       = "...my_request_signing_certificate..."
    request_signing_key               = "...my_request_signing_key..."
    response_digest_algorithm         = "SHA1"
    response_encryption_key           = "...my_response_encryption_key..."
    response_signature_algorithm      = "SHA384"
    session_absolute_timeout          = 8.17
    session_audience                  = "...my_session_audience..."
    session_cookie_domain             = "...my_session_cookie_domain..."
    session_cookie_http_only          = true
    session_cookie_name               = "...my_session_cookie_name..."
    session_cookie_path               = "...my_session_cookie_path..."
    session_cookie_same_site          = "None"
    session_cookie_secure             = true
    session_enforce_same_subject      = true
    session_hash_storage_key          = false
    session_hash_subject              = false
    session_idling_timeout            = 3.44
    session_memcached_host            = "...my_session_memcached_host..."
    session_memcached_port            = 59429
    session_memcached_prefix          = "...my_session_memcached_prefix..."
    session_memcached_socket          = "...my_session_memcached_socket..."
    session_remember                  = false
    session_remember_absolute_timeout = 4.84
    session_remember_cookie_name      = "...my_session_remember_cookie_name..."
    session_remember_rolling_timeout  = 7.93
    session_request_headers = [
      "id"
    ]
    session_response_headers = [
      "id"
    ]
    session_rolling_timeout      = 5.35
    session_secret               = "...my_session_secret..."
    session_storage              = "cookie"
    session_store_metadata       = false
    validate_assertion_signature = true
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
    "grpc"
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