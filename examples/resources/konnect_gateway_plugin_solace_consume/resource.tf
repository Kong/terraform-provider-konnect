resource "konnect_gateway_plugin_solace_consume" "my_gatewaypluginsolaceconsume" {
  config = {
    flow = {
      ack_mode = "CLIENT"
      binds = [
        {
          name = "...my_name..."
          type = "QUEUE"
        }
      ]
      functions = [
        "..."
      ]
      max_unacked_messages = 5
      properties = {
        key = "value"
      }
      selector     = "...my_selector..."
      wait_timeout = 2591
      window_size  = 141
    }
    mode = "WEBSOCKET"
    polling = {
      timeout = 287082
    }
    session = {
      authentication = {
        access_token        = "...my_access_token..."
        access_token_header = "...my_access_token_header..."
        basic_auth_header   = "...my_basic_auth_header..."
        id_token            = "...my_id_token..."
        id_token_header     = "...my_id_token_header..."
        password            = "...my_password..."
        scheme              = "BASIC"
        username            = "...my_username..."
      }
      calculate_message_expiry = false
      connect_timeout          = 27609
      generate_rcv_timestamps  = true
      generate_send_timestamps = true
      generate_sender_id       = false
      generate_sequence_number = true
      host                     = "...my_host..."
      properties = {
        key = "value"
      }
      ssl_validate_certificate = true
      vpn_name                 = "...my_vpn_name..."
    }
    websocket = {
      max_recv_len = 5
      max_send_len = 10
      timeout      = 44172
    }
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 8
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
  updated_at = 4
}