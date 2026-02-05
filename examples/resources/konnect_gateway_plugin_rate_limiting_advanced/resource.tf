resource "konnect_gateway_plugin_rate_limiting_advanced" "my_gatewaypluginratelimitingadvanced" {
  config = {
    compound_identifier = [
      "header"
    ]
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
    identifier              = "consumer"
    limit = [
      4.52
    ]
    lock_dictionary_name = "...my_lock_dictionary_name..."
    namespace            = "...my_namespace..."
    path                 = "...my_path..."
    redis = {
      cloud_authentication = {
        auth_provider            = "azure"
        aws_access_key_id        = "...my_aws_access_key_id..."
        aws_assume_role_arn      = "...my_aws_assume_role_arn..."
        aws_cache_name           = "...my_aws_cache_name..."
        aws_is_serverless        = true
        aws_region               = "...my_aws_region..."
        aws_role_session_name    = "...my_aws_role_session_name..."
        aws_secret_access_key    = "...my_aws_secret_access_key..."
        azure_client_id          = "...my_azure_client_id..."
        azure_client_secret      = "...my_azure_client_secret..."
        azure_tenant_id          = "...my_azure_tenant_id..."
        gcp_service_account_json = "...my_gcp_service_account_json..."
      }
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
      redis_proxy_type      = "envoy_v1.31"
      send_timeout          = 1710404950
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 57176
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "any"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      ssl               = true
      ssl_verify        = true
      username          = "...my_username..."
    }
    retry_after_jitter_max = 9.23
    strategy               = "cluster"
    sync_rate              = 4.11
    throttling = {
      enabled     = false
      interval    = 248813.69
      queue_limit = 503655.16
      retry_times = 989778.57
    }
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
  created_at       = 4
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
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
  protocols = [
    "grpcs"
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
  updated_at = 9
}