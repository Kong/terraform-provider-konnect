resource "konnect_gateway_plugin_confluent" "my_gatewaypluginconfluent" {
  config = {
    allowed_topics = [
      "..."
    ]
    bootstrap_servers = [
      {
        host = "...my_host..."
        port = 53037
      }
    ]
    cluster_api_key            = "...my_cluster_api_key..."
    cluster_api_secret         = "...my_cluster_api_secret..."
    cluster_name               = "...my_cluster_name..."
    confluent_cloud_api_key    = "...my_confluent_cloud_api_key..."
    confluent_cloud_api_secret = "...my_confluent_cloud_api_secret..."
    forward_body               = false
    forward_headers            = false
    forward_method             = false
    forward_uri                = true
    keepalive                  = 3
    keepalive_enabled          = true
    key_query_arg              = "...my_key_query_arg..."
    message_by_lua_functions = [
      "..."
    ]
    producer_async                                     = false
    producer_async_buffering_limits_messages_in_memory = 9
    producer_async_flush_timeout                       = 1
    producer_request_acks                              = 0
    producer_request_limits_bytes_per_request          = 7
    producer_request_limits_messages_per_request       = 6
    producer_request_retries_backoff_timeout           = 4
    producer_request_retries_max_attempts              = 2
    producer_request_timeout                           = 1
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
            auth_method               = "none"
            client_secret_jwt_alg     = "HS512"
            http_proxy                = "...my_http_proxy..."
            http_proxy_authorization  = "...my_http_proxy_authorization..."
            http_version              = 1.18
            https_proxy               = "...my_https_proxy..."
            https_proxy_authorization = "...my_https_proxy_authorization..."
            keep_alive                = false
            no_proxy                  = "...my_no_proxy..."
            ssl_verify                = true
            timeout                   = 170470379
          }
        }
        key_schema = {
          schema_version = "...my_schema_version..."
          subject_name   = "...my_subject_name..."
        }
        ssl_verify = true
        ttl        = 1001.01
        url        = "...my_url..."
        value_schema = {
          schema_version = "...my_schema_version..."
          subject_name   = "...my_subject_name..."
        }
      }
    }
    security = {
      ssl_verify = true
    }
    timeout          = 0
    topic            = "...my_topic..."
    topics_query_arg = "...my_topics_query_arg..."
  }
  consumer = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 0
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
  updated_at = 1
}