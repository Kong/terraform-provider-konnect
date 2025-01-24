resource "konnect_gateway_plugin_ai_request_transformer" "my_gatewaypluginairequesttransformer" {
  config = {
    http_proxy_host  = "...my_http_proxy_host..."
    http_proxy_port  = 19860
    http_timeout     = 10
    https_proxy_host = "...my_https_proxy_host..."
    https_proxy_port = 20590
    https_verify     = false
    llm = {
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
      logging = {
        log_payloads   = false
        log_statistics = false
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
          input_cost     = 6.37
          llama2_format  = "raw"
          max_tokens     = 5
          mistral_format = "openai"
          output_cost    = 8.25
          temperature    = 0.7
          top_k          = 420
          top_p          = 0.54
          upstream_path  = "...my_upstream_path..."
          upstream_url   = "...my_upstream_url..."
        }
        provider = "gemini"
      }
      route_type = "preserve"
    }
    max_request_body_size          = 7
    prompt                         = "...my_prompt..."
    transformation_extract_pattern = "...my_transformation_extract_pattern..."
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
    "tls_passthrough"
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