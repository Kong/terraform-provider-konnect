resource "konnect_gateway_plugin_kafka_log" "my_gatewaypluginkafkalog" {
  config = {
    authentication = {
      mechanism = "SCRAM-SHA-256"
      password  = "...my_password..."
      strategy  = "sasl"
      tokenauth = false
      user      = "...my_user..."
    }
    bootstrap_servers = [
      {
        host = "...my_host..."
        port = 7302
      }
    ]
    cluster_name = "...my_cluster_name..."
    custom_fields_by_lua = {
      key = "value"
    }
    keepalive                                          = 2
    keepalive_enabled                                  = true
    key_query_arg                                      = "...my_key_query_arg..."
    producer_async                                     = false
    producer_async_buffering_limits_messages_in_memory = 2
    producer_async_flush_timeout                       = 6
    producer_request_acks                              = 0
    producer_request_limits_bytes_per_request          = 5
    producer_request_limits_messages_per_request       = 10
    producer_request_retries_backoff_timeout           = 1
    producer_request_retries_max_attempts              = 4
    producer_request_timeout                           = 3
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
            auth_method               = "client_secret_jwt"
            client_secret_jwt_alg     = "HS512"
            http_proxy                = "...my_http_proxy..."
            http_proxy_authorization  = "...my_http_proxy_authorization..."
            http_version              = 9.06
            https_proxy               = "...my_https_proxy..."
            https_proxy_authorization = "...my_https_proxy_authorization..."
            keep_alive                = true
            no_proxy                  = "...my_no_proxy..."
            ssl_verify                = true
            timeout                   = 893347002
          }
        }
        key_schema = {
          schema_version = "...my_schema_version..."
          subject_name   = "...my_subject_name..."
        }
        ssl_verify = true
        ttl        = 94.95
        url        = "...my_url..."
        value_schema = {
          schema_version = "...my_schema_version..."
          subject_name   = "...my_subject_name..."
        }
      }
    }
    security = {
      certificate_id = "...my_certificate_id..."
      ssl            = false
      ssl_verify     = false
    }
    timeout = 2
    topic   = "...my_topic..."
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
  updated_at = 5
}