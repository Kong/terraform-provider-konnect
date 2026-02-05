resource "konnect_gateway_plugin_ace" "my_gatewaypluginace" {
  config = {
    anonymous    = "...my_anonymous..."
    match_policy = "required"
    rate_limiting = {
      redis = {
        cloud_authentication = {
          auth_provider            = "azure"
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
        cluster_max_redirections = 10
        cluster_nodes = [
          {
            ip   = "...my_ip..."
            port = 57941
          }
        ]
        connect_timeout       = 1979030271
        connection_is_proxied = false
        database              = 7
        host                  = "...my_host..."
        keepalive_backlog     = 162019354
        keepalive_pool_size   = 1227973735
        password              = "...my_password..."
        port                  = 61563
        read_timeout          = 837560096
        send_timeout          = 1354140762
        sentinel_master       = "...my_sentinel_master..."
        sentinel_nodes = [
          {
            host = "...my_host..."
            port = 41842
          }
        ]
        sentinel_password = "...my_sentinel_password..."
        sentinel_role     = "any"
        sentinel_username = "...my_sentinel_username..."
        server_name       = "...my_server_name..."
        ssl               = true
        ssl_verify        = false
        username          = "...my_username..."
      }
      sync_rate = 161.65
    }
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 6
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
  updated_at = 7
}