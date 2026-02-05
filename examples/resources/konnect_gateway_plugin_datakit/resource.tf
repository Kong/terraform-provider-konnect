resource "konnect_gateway_plugin_datakit" "my_gatewayplugindatakit" {
  config = {
    debug = false
    nodes = [
      {
        xml_to_json = {
          attributes_block_name  = "...my_attributes_block_name..."
          attributes_name_prefix = "...my_attributes_name_prefix..."
          input                  = "...my_input..."
          name                   = "...my_name..."
          output                 = "...my_output..."
          recognize_type         = true
          text_as_property       = true
          text_block_name        = "...my_text_block_name..."
          xpath                  = "...my_xpath..."
        }
      }
    ]
    resources = {
      cache = {
        memory = {
          dictionary_name = "...my_dictionary_name..."
        }
        redis = {
          cloud_authentication = {
            auth_provider            = "gcp"
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
          cluster_max_redirections = 10
          cluster_nodes = [
            {
              ip   = "...my_ip..."
              port = 23392
            }
          ]
          connect_timeout       = 2043965363
          connection_is_proxied = true
          database              = 6
          host                  = "...my_host..."
          keepalive_backlog     = 1403771087
          keepalive_pool_size   = 371071403
          password              = "...my_password..."
          port                  = 3296
          read_timeout          = 1474836361
          send_timeout          = 1614607440
          sentinel_master       = "...my_sentinel_master..."
          sentinel_nodes = [
            {
              host = "...my_host..."
              port = 15355
            }
          ]
          sentinel_password = "...my_sentinel_password..."
          sentinel_role     = "master"
          sentinel_username = "...my_sentinel_username..."
          server_name       = "...my_server_name..."
          ssl               = false
          ssl_verify        = false
          username          = "...my_username..."
        }
        strategy = "memory"
      }
      vault = {
        key = "value"
      }
    }
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 4
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