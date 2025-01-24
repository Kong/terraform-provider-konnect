resource "konnect_gateway_plugin_rate_limiting_advanced" "my_gatewaypluginratelimitingadvanced" {
  config = {
    consumer_groups = [
      "..."
    ]
    dictionary_name         = "...my_dictionary_name..."
    disable_penalty         = false
    enforce_consumer_groups = false
    error_code              = 3.07
    error_message           = "...my_error_message..."
    header_name             = "...my_header_name..."
    hide_client_headers     = true
    identifier              = "ip"
    limit = [
      4.52
    ]
    namespace = "...my_namespace..."
    path      = "...my_path..."
    redis = {
      cluster_max_redirections = 7
      cluster_nodes = [
        {
          ip   = "...my_ip..."
          port = 54354
        }
      ]
      connect_timeout       = 374037696
      connection_is_proxied = false
      database              = 3
      host                  = "...my_host..."
      keepalive_backlog     = 788639936
      keepalive_pool_size   = 1875079515
      password              = "...my_password..."
      port                  = 63483
      read_timeout          = 382621324
      send_timeout          = 1710404950
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 57176
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "master"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      ssl               = true
      ssl_verify        = true
      username          = "...my_username..."
    }
    retry_after_jitter_max = 9.23
    strategy               = "cluster"
    sync_rate              = 4.11
    window_size = [
      4.61
    ]
    window_type = "sliding"
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