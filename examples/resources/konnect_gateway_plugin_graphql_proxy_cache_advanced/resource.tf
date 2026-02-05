resource "konnect_gateway_plugin_graphql_proxy_cache_advanced" "my_gatewayplugingraphqlproxycacheadvanced" {
  config = {
    bypass_on_err = true
    cache_ttl     = 7
    memory = {
      dictionary_name = "...my_dictionary_name..."
    }
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
      cluster_max_redirections = 8
      cluster_nodes = [
        {
          ip   = "...my_ip..."
          port = 438
        }
      ]
      connect_timeout       = 974295670
      connection_is_proxied = true
      database              = 0
      host                  = "...my_host..."
      keepalive_backlog     = 329763938
      keepalive_pool_size   = 353109135
      password              = "...my_password..."
      port                  = 47126
      read_timeout          = 1749476727
      send_timeout          = 2078135811
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 41845
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "slave"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      ssl               = true
      ssl_verify        = false
      username          = "...my_username..."
    }
    strategy = "memory"
    vary_headers = [
      "..."
    ]
  }
  consumer = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 6
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
  updated_at = 5
}