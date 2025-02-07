resource "konnect_gateway_plugin_service_protection" "my_gatewaypluginserviceprotection" {
  config = {
    dictionary_name     = "...my_dictionary_name..."
    disable_penalty     = true
    error_code          = 2.21
    error_message       = "...my_error_message..."
    hide_client_headers = false
    limit = [
      3.19
    ]
    lock_dictionary_name = "...my_lock_dictionary_name..."
    namespace            = "...my_namespace..."
    redis = {
      cluster_max_redirections = 3
      cluster_nodes = [
        {
          ip   = "...my_ip..."
          port = 33693
        }
      ]
      connect_timeout       = 498772011
      connection_is_proxied = false
      database              = 3
      host                  = "...my_host..."
      keepalive_backlog     = 312843254
      keepalive_pool_size   = 1451118259
      password              = "...my_password..."
      port                  = 56514
      read_timeout          = 206958009
      send_timeout          = 408106576
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 32122
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "any"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      ssl               = false
      ssl_verify        = true
      username          = "...my_username..."
    }
    retry_after_jitter_max = 4.23
    strategy               = "redis"
    sync_rate              = 5.69
    window_size = [
      5.81
    ]
    window_type = "sliding"
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
    "http"
  ]
  service = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
}