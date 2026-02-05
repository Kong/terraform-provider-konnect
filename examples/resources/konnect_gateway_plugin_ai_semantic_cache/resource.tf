resource "konnect_gateway_plugin_ai_semantic_cache" "my_gatewaypluginaisemanticcache" {
  config = {
    cache_control = true
    cache_ttl     = 7
    embeddings = {
      auth = {
        allow_override             = false
        aws_access_key_id          = "...my_aws_access_key_id..."
        aws_secret_access_key      = "...my_aws_secret_access_key..."
        azure_client_id            = "...my_azure_client_id..."
        azure_client_secret        = "...my_azure_client_secret..."
        azure_tenant_id            = "...my_azure_tenant_id..."
        azure_use_managed_identity = true
        gcp_service_account_json   = "...my_gcp_service_account_json..."
        gcp_use_service_account    = true
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
            use_cache      = true
            wait_for_model = true
          }
          upstream_url = "...my_upstream_url..."
        }
        provider = "bedrock"
      }
    }
    exact_caching            = true
    ignore_assistant_prompts = true
    ignore_system_prompts    = false
    ignore_tool_prompts      = true
    llm_format               = "huggingface"
    message_countback        = 289.61
    stop_on_failure          = false
    vectordb = {
      dimensions      = 9
      distance_metric = "euclidean"
      pgvector = {
        database     = "...my_database..."
        host         = "...my_host..."
        password     = "...my_password..."
        port         = 4
        ssl          = false
        ssl_cert     = "...my_ssl_cert..."
        ssl_cert_key = "...my_ssl_cert_key..."
        ssl_required = false
        ssl_verify   = true
        ssl_version  = "any"
        timeout      = 6.41
        user         = "...my_user..."
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
        cluster_max_redirections = 8
        cluster_nodes = [
          {
            ip   = "...my_ip..."
            port = 18507
          }
        ]
        connect_timeout       = 282426182
        connection_is_proxied = false
        database              = 0
        host                  = "...my_host..."
        keepalive_backlog     = 673476807
        keepalive_pool_size   = 230676668
        password              = "...my_password..."
        port                  = 51340
        read_timeout          = 289855833
        send_timeout          = 1337575450
        sentinel_master       = "...my_sentinel_master..."
        sentinel_nodes = [
          {
            host = "...my_host..."
            port = 38845
          }
        ]
        sentinel_password = "...my_sentinel_password..."
        sentinel_role     = "slave"
        sentinel_username = "...my_sentinel_username..."
        server_name       = "...my_server_name..."
        ssl               = true
        ssl_verify        = false
        username          = "...my_username..."
      }
      strategy  = "redis"
      threshold = 5.91
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
  updated_at = 8
}