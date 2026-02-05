resource "konnect_gateway_plugin_ai_mcp_proxy" "my_gatewaypluginaimcpproxy" {
  config = {
    consumer_identifier = "custom_id"
    default_acl = [
      {
        allow = [
          "..."
        ]
        deny = [
          "..."
        ]
        scope = "...my_scope..."
      }
    ]
    include_consumer_groups = false
    logging = {
      log_audits     = true
      log_payloads   = false
      log_statistics = true
    }
    max_request_body_size = 8
    mode                  = "passthrough-listener"
    server = {
      forward_client_headers = true
      tag                    = "...my_tag..."
      timeout                = 7.85
    }
    tools = [
      {
        acl = {
          allow = [
            "..."
          ]
          deny = [
            "..."
          ]
        }
        annotations = {
          destructive_hint = true
          idempotent_hint  = false
          open_world_hint  = false
          read_only_hint   = false
          title            = "...my_title..."
        }
        description = "...my_description..."
        headers = {
          key = [
            # ...
          ]
        }
        host   = "...my_host..."
        method = "GET"
        name   = "...my_name..."
        parameters = [
          {
            additional_properties = "{ \"see\": \"documentation\" }"
            description           = "...my_description..."
            in                    = "...my_in..."
            name                  = "...my_name..."
            required              = true
            schema = {
              type = "...my_type..."
            }
          }
        ]
        path = "...my_path..."
        query = {
          key = [
            # ...
          ]
        }
        request_body = {
          key = jsonencode("value")
        }
        responses = {
          key = jsonencode("value")
        }
        scheme = "https"
      }
    ]
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 4
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
  updated_at = 1
}