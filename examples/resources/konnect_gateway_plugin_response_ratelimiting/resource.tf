resource "konnect_gateway_plugin_response_ratelimiting" "my_gatewaypluginresponseratelimiting" {
  config = {
    block_on_first_violation = true
    fault_tolerant           = false
    header_name              = "...my_header_name..."
    hide_client_headers      = true
    limit_by                 = "ip"
    limits = {
      key = {
        day    = 6.02
        hour   = 1.45
        minute = 3.9
        month  = 1.45
        second = 8.28
        year   = 4.52
      }
    }
    policy = "cluster"
    redis = {
      cloud_authentication = {
        auth_provider            = "gcp"
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
      database    = 9
      host        = "...my_host..."
      password    = "...my_password..."
      port        = 32565
      server_name = "...my_server_name..."
      ssl         = true
      ssl_verify  = false
      timeout     = 70842937
      username    = "...my_username..."
    }
  }
  consumer = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 10
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
  updated_at = 5
}