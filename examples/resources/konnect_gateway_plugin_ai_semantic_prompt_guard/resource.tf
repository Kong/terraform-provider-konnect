resource "konnect_gateway_plugin_ai_semantic_prompt_guard" "my_gatewaypluginaisemanticpromptguard" {
  config = {
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
        param_location             = "query"
        param_name                 = "...my_param_name..."
        param_value                = "...my_param_value..."
      }
      model = {
        name = "...my_name..."
        options = {
          upstream_url = "...my_upstream_url..."
        }
        provider = "openai"
      }
    }
    rules = {
      allow_prompts = [
        "..."
      ]
      deny_prompts = [
        "..."
      ]
      match_all_conversation_history = false
      match_all_roles                = false
      max_request_body_size          = 10
    }
    search = {
      threshold = 9.87
    }
    vectordb = {
      dimensions      = 4
      distance_metric = "euclidean"
      redis = {
        cluster_max_redirections = 10
        cluster_nodes = [
          {
            ip   = "...my_ip..."
            port = 4694
          }
        ]
        connect_timeout       = 927722545
        connection_is_proxied = false
        database              = 7
        host                  = "...my_host..."
        keepalive_backlog     = 1501125607
        keepalive_pool_size   = 114222283
        password              = "...my_password..."
        port                  = 24350
        read_timeout          = 888592257
        send_timeout          = 228434190
        sentinel_master       = "...my_sentinel_master..."
        sentinel_nodes = [
          {
            host = "...my_host..."
            port = 18183
          }
        ]
        sentinel_password = "...my_sentinel_password..."
        sentinel_role     = "any"
        sentinel_username = "...my_sentinel_username..."
        server_name       = "...my_server_name..."
        ssl               = false
        ssl_verify        = true
        username          = "...my_username..."
      }
      strategy  = "redis"
      threshold = 3.9
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
}