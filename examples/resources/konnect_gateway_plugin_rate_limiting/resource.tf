resource "konnect_gateway_plugin_rate_limiting" "my_gatewaypluginratelimiting" {
  config = {
    day                 = 2.06
    error_code          = 7.45
    error_message       = "...my_error_message..."
    fault_tolerant      = true
    header_name         = "...my_header_name..."
    hide_client_headers = false
    hour                = 0.92
    limit_by            = "header"
    minute              = 3.29
    month               = 7.22
    path                = "...my_path..."
    policy              = "cluster"
    redis = {
      database    = 3
      host        = "...my_host..."
      password    = "...my_password..."
      port        = 37089
      server_name = "...my_server_name..."
      ssl         = false
      ssl_verify  = true
      timeout     = 14058328
      username    = "...my_username..."
    }
    second    = 1.89
    sync_rate = 7.98
    year      = 0.32
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
}