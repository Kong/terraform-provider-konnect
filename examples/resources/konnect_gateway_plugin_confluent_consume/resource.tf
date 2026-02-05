resource "konnect_gateway_plugin_confluent_consume" "my_gatewaypluginconfluentconsume" {
  config = {
    auto_offset_reset = "earliest"
    bootstrap_servers = [
      {
        host = "...my_host..."
        port = 58521
      }
    ]
    cluster_api_key             = "...my_cluster_api_key..."
    cluster_api_secret          = "...my_cluster_api_secret..."
    cluster_name                = "...my_cluster_name..."
    commit_strategy             = "auto"
    confluent_cloud_api_key     = "...my_confluent_cloud_api_key..."
    confluent_cloud_api_secret  = "...my_confluent_cloud_api_secret..."
    dlq_topic                   = "...my_dlq_topic..."
    enable_dlq                  = true
    enforce_latest_offset_reset = false
    keepalive                   = 2
    keepalive_enabled           = true
    message_by_lua_functions = [
      "..."
    ]
    message_deserializer = "noop"
    mode                 = "http-get"
    schema_registry = {
      confluent = {
        authentication = {
          basic = {
            password = "...my_password..."
            username = "...my_username..."
          }
          mode = "basic"
          oauth2 = {
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
          oauth2_client = {
            auth_method               = "client_secret_basic"
            client_secret_jwt_alg     = "HS512"
            http_proxy                = "...my_http_proxy..."
            http_proxy_authorization  = "...my_http_proxy_authorization..."
            http_version              = 8.6
            https_proxy               = "...my_https_proxy..."
            https_proxy_authorization = "...my_https_proxy_authorization..."
            keep_alive                = false
            no_proxy                  = "...my_no_proxy..."
            ssl_verify                = true
            timeout                   = 595962925
          }
        }
        ssl_verify = false
        ttl        = 2309.04
        url        = "...my_url..."
      }
    }
    security = {
      ssl_verify = false
    }
    timeout = 7
    topics = [
      {
        name = "...my_name..."
        schema_registry = {
          confluent = {
            authentication = {
              basic = {
                password = "...my_password..."
                username = "...my_username..."
              }
              mode = "basic"
              oauth2 = {
                audience = [
                  "..."
                ]
                client_id     = "...my_client_id..."
                client_secret = "...my_client_secret..."
                grant_type    = "client_credentials"
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
              oauth2_client = {
                auth_method               = "client_secret_basic"
                client_secret_jwt_alg     = "HS512"
                http_proxy                = "...my_http_proxy..."
                http_proxy_authorization  = "...my_http_proxy_authorization..."
                http_version              = 5.07
                https_proxy               = "...my_https_proxy..."
                https_proxy_authorization = "...my_https_proxy_authorization..."
                keep_alive                = false
                no_proxy                  = "...my_no_proxy..."
                ssl_verify                = true
                timeout                   = 1747283445
              }
            }
            ssl_verify = true
            ttl        = 179.01
            url        = "...my_url..."
          }
        }
      }
    ]
  }
  consumer = {
    id = "...my_id..."
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
    "wss"
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
  updated_at = 0
}