resource "konnect_gateway_plugin_basic_auth" "my_gatewaypluginbasicauth" {
  config = {
    anonymous = "...my_anonymous..."
    brute_force_protection = {
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
        database    = 2
        host        = "...my_host..."
        password    = "...my_password..."
        port        = 18528
        server_name = "...my_server_name..."
        ssl         = true
        ssl_verify  = true
        timeout     = 1835674936
        username    = "...my_username..."
      }
      strategy = "redis"
    }
    hide_credentials = false
    realm            = "...my_realm..."
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
  updated_at = 10
}