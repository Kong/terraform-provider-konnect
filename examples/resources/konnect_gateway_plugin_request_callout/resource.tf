resource "konnect_gateway_plugin_request_callout" "my_gatewaypluginrequestcallout" {
  config = {
    cache = {
      cache_ttl = 7
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
            port = 12116
          }
        ]
        connect_timeout       = 221492024
        connection_is_proxied = false
        database              = 0
        host                  = "...my_host..."
        keepalive_backlog     = 2070085081
        keepalive_pool_size   = 254346691
        password              = "...my_password..."
        port                  = 2887
        read_timeout          = 1913162680
        send_timeout          = 27958014
        sentinel_master       = "...my_sentinel_master..."
        sentinel_nodes = [
          {
            host = "...my_host..."
            port = 48331
          }
        ]
        sentinel_password = "...my_sentinel_password..."
        sentinel_role     = "master"
        sentinel_username = "...my_sentinel_username..."
        server_name       = "...my_server_name..."
        ssl               = true
        ssl_verify        = false
        username          = "...my_username..."
      }
      strategy = "redis"
    }
    callouts = [
      {
        cache = {
          bypass = true
        }
        depends_on = [
          "..."
        ]
        name = "...my_name..."
        request = {
          body = {
            custom = {
              key = "value"
            }
            decode  = true
            forward = false
          }
          by_lua = "...my_by_lua..."
          error = {
            error_response_code = 5
            error_response_msg  = "...my_error_response_msg..."
            http_statuses = [
              791
            ]
            on_error = "fail"
            retries  = 4
          }
          headers = {
            custom = {
              key = "value"
            }
            forward = false
          }
          http_opts = {
            proxy = {
              auth_password = "...my_auth_password..."
              auth_username = "...my_auth_username..."
              http_proxy    = "...my_http_proxy..."
              https_proxy   = "...my_https_proxy..."
            }
            ssl_server_name = "...my_ssl_server_name..."
            ssl_verify      = true
            timeouts = {
              connect = 396369045
              read    = 2097982023
              write   = 1346258560
            }
          }
          method = "...my_method..."
          query = {
            custom = {
              key = "value"
            }
            forward = true
          }
          url = "...my_url..."
        }
        response = {
          body = {
            decode = false
            store  = true
          }
          by_lua = "...my_by_lua..."
          headers = {
            store = false
          }
        }
      }
    ]
    upstream = {
      body = {
        custom = {
          key = "value"
        }
        decode  = true
        forward = false
      }
      by_lua = "...my_by_lua..."
      headers = {
        custom = {
          key = "value"
        }
        forward = true
      }
      query = {
        custom = {
          key = "value"
        }
        forward = false
      }
    }
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
  updated_at = 1
}