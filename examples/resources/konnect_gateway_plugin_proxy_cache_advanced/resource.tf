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
      cloud_authentication = {
        auth_provider            = "aws"
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
      sentinel_role     = "any"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      ssl               = false
      ssl_verify        = true
      username          = "...my_username..."
    }
    request_method = [
      "GET"
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
  created_at       = 8
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
  updated_at = 6
}