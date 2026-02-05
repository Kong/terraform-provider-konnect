resource "konnect_gateway_partial" "my_gatewaypartial" {
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  redis_ce = {
    config = {
      cloud_authentication = {
        auth_provider            = "gcp"
        aws_access_key_id        = "...my_aws_access_key_id..."
        aws_assume_role_arn      = "...my_aws_assume_role_arn..."
        aws_cache_name           = "...my_aws_cache_name..."
        aws_is_serverless        = false
        aws_region               = "...my_aws_region..."
        aws_role_session_name    = "...my_aws_role_session_name..."
        aws_secret_access_key    = "...my_aws_secret_access_key..."
        azure_client_id          = "...my_azure_client_id..."
        azure_client_secret      = "...my_azure_client_secret..."
        azure_tenant_id          = "...my_azure_tenant_id..."
        gcp_service_account_json = "...my_gcp_service_account_json..."
      }
      database    = 5
      host        = "...my_host..."
      password    = "...my_password..."
      port        = 8055
      server_name = "...my_server_name..."
      ssl         = true
      ssl_verify  = true
      timeout     = 2089324356
      username    = "...my_username..."
    }
    created_at = 10
    id         = "...my_id..."
    name       = "...my_name..."
    tags = [
      "..."
    ]
    updated_at = 2
  }
  redis_ee = {
    config = {
      cloud_authentication = {
        auth_provider            = "gcp"
        aws_access_key_id        = "...my_aws_access_key_id..."
        aws_assume_role_arn      = "...my_aws_assume_role_arn..."
        aws_cache_name           = "...my_aws_cache_name..."
        aws_is_serverless        = false
        aws_region               = "...my_aws_region..."
        aws_role_session_name    = "...my_aws_role_session_name..."
        aws_secret_access_key    = "...my_aws_secret_access_key..."
        azure_client_id          = "...my_azure_client_id..."
        azure_client_secret      = "...my_azure_client_secret..."
        azure_tenant_id          = "...my_azure_tenant_id..."
        gcp_service_account_json = "...my_gcp_service_account_json..."
      }
      cluster_max_redirections = 4
      cluster_nodes = [
        {
          ip   = "...my_ip..."
          port = 13662
        }
      ]
      connect_timeout       = 1346533360
      connection_is_proxied = false
      database              = 1
      host                  = "...my_host..."
      keepalive_backlog     = 259647341
      keepalive_pool_size   = 1491672736
      password              = "...my_password..."
      port                  = 17730
      read_timeout          = 1431389868
      send_timeout          = 1316344582
      sentinel_master       = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "...my_host..."
          port = 32792
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "slave"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      ssl               = false
      ssl_verify        = false
      username          = "...my_username..."
    }
    created_at = 2
    id         = "...my_id..."
    name       = "...my_name..."
    tags = [
      "..."
    ]
    updated_at = 8
  }
}