resource "konnect_gateway_plugin_upstream_oauth" "my_gatewaypluginupstreamoauth" {
  config = {
    behavior = {
      idp_error_response_body_template = "...my_idp_error_response_body_template..."
      idp_error_response_content_type  = "...my_idp_error_response_content_type..."
      idp_error_response_message       = "...my_idp_error_response_message..."
      idp_error_response_status_code   = 576
      purge_token_on_upstream_status_codes = [
        373
      ]
      upstream_access_token_header_name = "...my_upstream_access_token_header_name..."
    }
    cache = {
      default_ttl    = 7.94
      eagerly_expire = 9
      memory = {
        dictionary_name = "...my_dictionary_name..."
      }
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
        cluster_max_redirections = 10
        cluster_nodes = [
          {
            ip   = "...my_ip..."
            port = 20643
          }
        ]
        connect_timeout       = 305131733
        connection_is_proxied = true
        database              = 10
        host                  = "...my_host..."
        keepalive_backlog     = 1047987263
        keepalive_pool_size   = 459234090
        password              = "...my_password..."
        port                  = 35119
        read_timeout          = 245223357
        send_timeout          = 1142057358
        sentinel_master       = "...my_sentinel_master..."
        sentinel_nodes = [
          {
            host = "...my_host..."
            port = 31719
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
      strategy = "memory"
    }
    client = {
      auth_method               = "none"
      client_secret_jwt_alg     = "HS256"
      http_proxy                = "...my_http_proxy..."
      http_proxy_authorization  = "...my_http_proxy_authorization..."
      http_version              = 6.12
      https_proxy               = "...my_https_proxy..."
      https_proxy_authorization = "...my_https_proxy_authorization..."
      keep_alive                = true
      no_proxy                  = "...my_no_proxy..."
      ssl_verify                = false
      timeout                   = 1421616738
    }
    oauth = {
      audience = [
        "..."
      ]
      client_id     = "...my_client_id..."
      client_secret = "...my_client_secret..."
      grant_type    = "password"
      password      = "...my_password..."
      scopes = [
        "..."
      ]
      token_endpoint = "...my_token_endpoint..."
      token_headers = {
        key = "value"
      }
      token_post_args = {
        key = "value"
      }
      username = "...my_username..."
    }
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 5
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