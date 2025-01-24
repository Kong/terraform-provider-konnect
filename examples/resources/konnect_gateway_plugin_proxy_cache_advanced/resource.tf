resource "konnect_gateway_plugin_proxy_cache_advanced" "my_gatewaypluginproxycacheadvanced" {
  config = {
    bypass_on_err = true
    cache_control = true
    cache_ttl     = 5
    content_type = [
      "..."
    ]
    ignore_uri_case = false
    memory = {
      dictionary_name = "...my_dictionary_name..."
    }
    redis = {
      cluster_max_redirections = 7
      cluster_nodes = [
        {
          ip   = "...my_ip..."
          port = 56343
        }
      ]
      connect_timeout       = 883264270
      connection_is_proxied = false
      database              = 1
      host                  = "...my_host..."
      keepalive_backlog     = 578209368
      keepalive_pool_size   = 1307431457
      password              = "...my_password..."
      port                  = 54281
      read_timeout          = 350076819
      send_timeout          = 2140614627
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 8222
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "master"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      ssl               = false
      ssl_verify        = true
      username          = "...my_username..."
    }
    request_method = [
      "HEAD"
    ]
    response_code = [
      269
    ]
    response_headers = {
      age            = false
      x_cache_key    = true
      x_cache_status = true
    }
    storage_ttl = 0
    strategy    = "redis"
    vary_headers = [
      "..."
    ]
    vary_query_params = [
      "..."
    ]
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
    "http"
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