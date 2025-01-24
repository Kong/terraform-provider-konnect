resource "konnect_gateway_plugin_graphql_rate_limiting_advanced" "my_gatewayplugingraphqlratelimitingadvanced" {
  config = {
    cost_strategy       = "default"
    dictionary_name     = "...my_dictionary_name..."
    hide_client_headers = true
    identifier          = "credential"
    limit = [
      1.18
    ]
    max_cost  = 5.02
    namespace = "...my_namespace..."
    redis = {
      cluster_max_redirections = 0
      cluster_nodes = [
        {
          ip   = "...my_ip..."
          port = 11194
        }
      ]
      connect_timeout       = 374060135
      connection_is_proxied = false
      database              = 0
      host                  = "...my_host..."
      keepalive_backlog     = 171042224
      keepalive_pool_size   = 2132834450
      password              = "...my_password..."
      port                  = 18586
      read_timeout          = 329378340
      send_timeout          = 778888138
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 63567
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "master"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      ssl               = true
      ssl_verify        = false
      username          = "...my_username..."
    }
    score_factor = 5.13
    strategy     = "redis"
    sync_rate    = 5.2
    window_size = [
      1.07
    ]
    window_type = "fixed"
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
    "udp"
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