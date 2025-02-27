resource "konnect_gateway_plugin_ai_rate_limiting_advanced" "my_gatewaypluginairatelimitingadvanced" {
  config = {
    dictionary_name      = "...my_dictionary_name..."
    disable_penalty      = false
    error_code           = 8.18
    error_hide_providers = false
    error_message        = "...my_error_message..."
    header_name          = "...my_header_name..."
    hide_client_headers  = false
    identifier           = "credential"
    llm_providers = [
      {
        limit       = 1.4
        name        = "huggingface"
        window_size = 6.05
      }
    ]
    path = "...my_path..."
    redis = {
      cluster_max_redirections = 5
      cluster_nodes = [
        {
          ip   = "...my_ip..."
          port = 36904
        }
      ]
      connect_timeout       = 883187098
      connection_is_proxied = true
      database              = 5
      host                  = "...my_host..."
      keepalive_backlog     = 125882903
      keepalive_pool_size   = 165809619
      password              = "...my_password..."
      port                  = 37236
      read_timeout          = 394887043
      send_timeout          = 590662462
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 35255
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
    request_prompt_count_function = "...my_request_prompt_count_function..."
    retry_after_jitter_max        = 8.14
    strategy                      = "redis"
    sync_rate                     = 3.98
    tokens_count_strategy         = "prompt_tokens"
    window_type                   = "sliding"
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