resource "konnect_gateway_plugin_solace_log" "my_gatewaypluginsolacelog" {
  config = {
    message = {
      ack_timeout = 49183
      custom_fields_by_lua = {
        key = "value"
      }
      delivery_mode = "PERSISTENT"
      destinations = [
        {
          name = "...my_name..."
          type = "QUEUE"
        }
      ]
      dmq_eligible    = false
      priority        = 228
      sender_id       = "...my_sender_id..."
      tracing         = false
      tracing_sampled = false
      ttl             = 10
    }
    session = {
      authentication = {
        access_token        = "...my_access_token..."
        access_token_header = "...my_access_token_header..."
        basic_auth_header   = "...my_basic_auth_header..."
        id_token            = "...my_id_token..."
        id_token_header     = "...my_id_token_header..."
        password            = "...my_password..."
        scheme              = "NONE"
        username            = "...my_username..."
      }
      calculate_message_expiry = false
      connect_timeout          = 98544
      generate_rcv_timestamps  = false
      generate_send_timestamps = false
      generate_sender_id       = true
      generate_sequence_number = false
      host                     = "...my_host..."
      properties = {
        key = "value"
      }
      ssl_validate_certificate = true
      vpn_name                 = "...my_vpn_name..."
    }
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 2
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
  updated_at = 4
}