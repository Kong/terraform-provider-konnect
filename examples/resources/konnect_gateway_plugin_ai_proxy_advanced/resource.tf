resource "konnect_gateway_plugin_ai_proxy_advanced" "my_gatewaypluginaiproxyadvanced" {
  config = {
    balancer = {
      algorithm             = "lowest-latency"
      connect_timeout       = 1069677678
      hash_on_header        = "...my_hash_on_header..."
      latency_strategy      = "e2e"
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
          upstream_url = "...my_upstream_url..."
        }
        provider = "mistral"
      }
    }
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
        model = {
          name = "...my_name..."
          options = {
            anthropic_version   = "...my_anthropic_version..."
            azure_api_version   = "...my_azure_api_version..."
            azure_deployment_id = "...my_azure_deployment_id..."
            azure_instance      = "...my_azure_instance..."
            bedrock = {
              aws_region = "...my_aws_region..."
            }
            gemini = {
              api_endpoint = "...my_api_endpoint..."
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
          provider = "bedrock"
        }
        route_type = "llm/v1/completions"
        weight     = 58189
      }
    ]
    vectordb = {
      dimensions      = 3
      distance_metric = "euclidean"
      redis = {
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
      strategy  = "redis"
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
  protocols = [
    "https"
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
}