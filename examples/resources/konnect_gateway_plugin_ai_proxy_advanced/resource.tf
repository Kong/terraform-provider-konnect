resource "konnect_gateway_plugin_ai_proxy_advanced" "my_gatewaypluginaiproxyadvanced" {
  config = {
    balancer = {
      algorithm       = "lowest-latency"
      connect_timeout = 1069677678
      fail_timeout    = 251767862
      failover_criteria = [
        "http_403"
      ]
      hash_on_header        = "...my_hash_on_header..."
      latency_strategy      = "e2e"
      max_fails             = 32489
      read_timeout          = 1128540479
      retries               = 14809
      slots                 = 27573
      tokens_count_strategy = "prompt-tokens"
      write_timeout         = 1475900303
    }
    embeddings = {
      auth = {
        allow_override             = true
        aws_access_key_id          = "...my_aws_access_key_id..."
        aws_secret_access_key      = "...my_aws_secret_access_key..."
        azure_client_id            = "...my_azure_client_id..."
        azure_client_secret        = "...my_azure_client_secret..."
        azure_tenant_id            = "...my_azure_tenant_id..."
        azure_use_managed_identity = false
        gcp_service_account_json   = "...my_gcp_service_account_json..."
        gcp_use_service_account    = false
        header_name                = "...my_header_name..."
        header_value               = "...my_header_value..."
        param_location             = "body"
        param_name                 = "...my_param_name..."
        param_value                = "...my_param_value..."
      }
      model = {
        name = "...my_name..."
        options = {
          azure = {
            api_version   = "...my_api_version..."
            deployment_id = "...my_deployment_id..."
            instance      = "...my_instance..."
          }
          bedrock = {
            aws_assume_role_arn        = "...my_aws_assume_role_arn..."
            aws_region                 = "...my_aws_region..."
            aws_role_session_name      = "...my_aws_role_session_name..."
            aws_sts_endpoint_url       = "...my_aws_sts_endpoint_url..."
            embeddings_normalize       = true
            performance_config_latency = "...my_performance_config_latency..."
            video_output_s3_uri        = "...my_video_output_s3_uri..."
          }
          gemini = {
            api_endpoint = "...my_api_endpoint..."
            location_id  = "...my_location_id..."
            project_id   = "...my_project_id..."
          }
          huggingface = {
            use_cache      = false
            wait_for_model = false
          }
          upstream_url = "...my_upstream_url..."
        }
        provider = "gemini"
      }
    }
    genai_category        = "realtime/generation"
    llm_format            = "huggingface"
    max_request_body_size = 5
    model_name_header     = true
    response_streaming    = "allow"
    targets = [
      {
        auth = {
          allow_override             = true
          aws_access_key_id          = "...my_aws_access_key_id..."
          aws_secret_access_key      = "...my_aws_secret_access_key..."
          azure_client_id            = "...my_azure_client_id..."
          azure_client_secret        = "...my_azure_client_secret..."
          azure_tenant_id            = "...my_azure_tenant_id..."
          azure_use_managed_identity = true
          gcp_service_account_json   = "...my_gcp_service_account_json..."
          gcp_use_service_account    = false
          header_name                = "...my_header_name..."
          header_value               = "...my_header_value..."
          param_location             = "query"
          param_name                 = "...my_param_name..."
          param_value                = "...my_param_value..."
        }
        description = "...my_description..."
        logging = {
          log_payloads   = true
          log_statistics = true
        }
        metadata = {
          key = jsonencode("value")
        }
        model = {
          name = "...my_name..."
          options = {
            anthropic_version   = "...my_anthropic_version..."
            azure_api_version   = "...my_azure_api_version..."
            azure_deployment_id = "...my_azure_deployment_id..."
            azure_instance      = "...my_azure_instance..."
            bedrock = {
              aws_assume_role_arn        = "...my_aws_assume_role_arn..."
              aws_region                 = "...my_aws_region..."
              aws_role_session_name      = "...my_aws_role_session_name..."
              aws_sts_endpoint_url       = "...my_aws_sts_endpoint_url..."
              embeddings_normalize       = true
              performance_config_latency = "...my_performance_config_latency..."
              video_output_s3_uri        = "...my_video_output_s3_uri..."
            }
            cohere = {
              embedding_input_type = "classification"
              wait_for_model       = true
            }
            dashscope = {
              international = true
            }
            embeddings_dimensions = 6
            gemini = {
              api_endpoint = "...my_api_endpoint..."
              endpoint_id  = "...my_endpoint_id..."
              location_id  = "...my_location_id..."
              project_id   = "...my_project_id..."
            }
            huggingface = {
              use_cache      = true
              wait_for_model = false
            }
            input_cost     = 2.57
            llama2_format  = "openai"
            max_tokens     = 2
            mistral_format = "openai"
            output_cost    = 7.34
            temperature    = 3.51
            top_k          = 204
            top_p          = 0.37
            upstream_path  = "...my_upstream_path..."
            upstream_url   = "...my_upstream_url..."
          }
          provider = "cerebras"
        }
        route_type = "llm/v1/assistants"
        weight     = 58189
      }
    ]
    vectordb = {
      dimensions      = 3
      distance_metric = "euclidean"
      pgvector = {
        database     = "...my_database..."
        host         = "...my_host..."
        password     = "...my_password..."
        port         = 5
        ssl          = true
        ssl_cert     = "...my_ssl_cert..."
        ssl_cert_key = "...my_ssl_cert_key..."
        ssl_required = true
        ssl_verify   = true
        ssl_version  = "any"
        timeout      = 3.81
        user         = "...my_user..."
      }
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
        cluster_max_redirections = 4
        cluster_nodes = [
          {
            ip   = "...my_ip..."
            port = 50944
          }
        ]
        connect_timeout       = 656443886
        connection_is_proxied = false
        database              = 10
        host                  = "...my_host..."
        keepalive_backlog     = 251172057
        keepalive_pool_size   = 1127137192
        password              = "...my_password..."
        port                  = 31201
        read_timeout          = 1222450418
        send_timeout          = 1541453227
        sentinel_master       = "...my_sentinel_master..."
        sentinel_nodes = [
          {
            host = "...my_host..."
            port = 61553
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
      strategy  = "pgvector"
      threshold = 6.56
    }
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 3
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
    "ws"
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
  updated_at = 8
}