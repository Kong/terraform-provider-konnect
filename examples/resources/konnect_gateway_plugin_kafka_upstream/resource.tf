resource "konnect_gateway_plugin_kafka_upstream" "my_gatewaypluginkafkaupstream" {
  config = {
    allowed_topics = [
      "..."
    ]
    authentication = {
      mechanism = "SCRAM-SHA-512"
      password  = "...my_password..."
      strategy  = "sasl"
      tokenauth = false
      user      = "...my_user..."
    }
    bootstrap_servers = [
      {
        host = "...my_host..."
        port = 3969
      }
    ]
    cluster_name      = "...my_cluster_name..."
    forward_body      = false
    forward_headers   = true
    forward_method    = true
    forward_uri       = false
    keepalive         = 3
    keepalive_enabled = true
    key_query_arg     = "...my_key_query_arg..."
    message_by_lua_functions = [
      "..."
    ]
    producer_async                                     = true
    producer_async_buffering_limits_messages_in_memory = 5
    producer_async_flush_timeout                       = 5
    producer_request_acks                              = 1
    producer_request_limits_bytes_per_request          = 1
    producer_request_limits_messages_per_request       = 9
    producer_request_retries_backoff_timeout           = 4
    producer_request_retries_max_attempts              = 5
    producer_request_timeout                           = 10
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
            auth_method               = "client_secret_jwt"
            client_secret_jwt_alg     = "HS512"
            http_proxy                = "...my_http_proxy..."
            http_proxy_authorization  = "...my_http_proxy_authorization..."
            http_version              = 7.51
            https_proxy               = "...my_https_proxy..."
            https_proxy_authorization = "...my_https_proxy_authorization..."
            keep_alive                = true
            no_proxy                  = "...my_no_proxy..."
            ssl_verify                = false
            timeout                   = 990649997
          }
        }
        key_schema = {
          schema_version = "...my_schema_version..."
          subject_name   = "...my_subject_name..."
        }
        ssl_verify = true
        ttl        = 1898.36
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
      ssl_verify     = true
    }
    timeout          = 0
    topic            = "...my_topic..."
    topics_query_arg = "...my_topics_query_arg..."
  }
  consumer = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 10
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
  updated_at = 0
}