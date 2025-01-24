resource "konnect_gateway_plugin_ai_proxy" "my_gatewaypluginaiproxy" {
  config = {
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
      param_location             = "body"
      param_name                 = "...my_param_name..."
      param_value                = "...my_param_value..."
    }
    logging = {
      log_payloads   = false
      log_statistics = true
    }
    max_request_body_size = 10
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
        input_cost     = 7.42
        llama2_format  = "openai"
        max_tokens     = 9
        mistral_format = "openai"
        output_cost    = 1.81
        temperature    = 2.26
        top_k          = 359
        top_p          = 0.14
        upstream_path  = "...my_upstream_path..."
        upstream_url   = "...my_upstream_url..."
      }
      provider = "openai"
    }
    model_name_header  = true
    response_streaming = "allow"
    route_type         = "llm/v1/chat"
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
    "tcp"
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