resource "konnect_gateway_plugin_response_ratelimiting" "my_gatewaypluginresponseratelimiting" {
  config = {
    block_on_first_violation = true
    fault_tolerant           = false
    header_name              = "...my_header_name..."
    hide_client_headers      = true
    limit_by                 = "ip"
    limits = {
      key = jsonencode("value"),
    }
    policy = "local"
    redis = {
      database    = 9
      host        = "...my_host..."
      password    = "...my_password..."
      port        = 32565
      server_name = "...my_server_name..."
      ssl         = true
      ssl_verify  = false
      timeout     = 70842937
      username    = "...my_username..."
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