resource "konnect_gateway_plugin_kafka_consume" "my_gatewaypluginkafkaconsume" {
  config = {
    authentication = {
      mechanism = "SCRAM-SHA-512"
      password  = "...my_password..."
      strategy  = "sasl"
      tokenauth = true
      user      = "...my_user..."
    }
    auto_offset_reset = "earliest"
    bootstrap_servers = [
      {
        host = "...my_host..."
        port = 27325
      }
    ]
    cluster_name                = "...my_cluster_name..."
    commit_strategy             = "off"
    dlq_topic                   = "...my_dlq_topic..."
    enable_dlq                  = false
    enforce_latest_offset_reset = false
    message_by_lua_functions = [
      "..."
    ]
    message_deserializer = "noop"
    mode                 = "websocket"
    schema_registry = {
      confluent = {
        authentication = {
          basic = {
            password = "...my_password..."
            username = "...my_username..."
          }
          mode = "oauth2"
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
            client_secret_jwt_alg     = "HS256"
            http_proxy                = "...my_http_proxy..."
            http_proxy_authorization  = "...my_http_proxy_authorization..."
            http_version              = 6.33
            https_proxy               = "...my_https_proxy..."
            https_proxy_authorization = "...my_https_proxy_authorization..."
            keep_alive                = false
            no_proxy                  = "...my_no_proxy..."
            ssl_verify                = true
            timeout                   = 2132166923
          }
        }
        ssl_verify = false
        ttl        = 1981.8
        url        = "...my_url..."
      }
    }
    security = {
      certificate_id = "...my_certificate_id..."
      ssl            = true
      ssl_verify     = false
    }
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
              mode = "none"
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
                client_secret_jwt_alg     = "HS256"
                http_proxy                = "...my_http_proxy..."
                http_proxy_authorization  = "...my_http_proxy_authorization..."
                http_version              = 3.98
                https_proxy               = "...my_https_proxy..."
                https_proxy_authorization = "...my_https_proxy_authorization..."
                keep_alive                = true
                no_proxy                  = "...my_no_proxy..."
                ssl_verify                = false
                timeout                   = 217890936
              }
            }
            ssl_verify = true
            ttl        = 2977.64
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
  created_at       = 0
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
    "wss"
  ]
  route = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
  updated_at = 0
}