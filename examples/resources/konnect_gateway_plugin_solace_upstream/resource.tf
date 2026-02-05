resource "konnect_gateway_plugin_solace_upstream" "my_gatewaypluginsolaceupstream" {
  config = {
    message = {
      ack_timeout     = 89103
      default_content = "...my_default_content..."
      delivery_mode   = "DIRECT"
      destinations = [
        {
          name = "...my_name..."
          type = "TOPIC"
        }
      ]
      dmq_eligible    = true
      forward_body    = false
      forward_headers = false
      forward_method  = true
      forward_uri     = false
      functions = [
        "..."
      ]
      priority        = 106
      sender_id       = "...my_sender_id..."
      tracing         = false
      tracing_sampled = false
      ttl             = 9
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
      connect_timeout          = 82788
      generate_rcv_timestamps  = false
      generate_send_timestamps = false
      generate_sender_id       = false
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
  created_at       = 5
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
  updated_at = 7
}