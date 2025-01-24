resource "konnect_gateway_plugin_statsd" "my_gatewaypluginstatsd" {
  config = {
    allow_status_codes = [
      "..."
    ]
    consumer_identifier_default = "username"
    flush_timeout               = 5.68
    host                        = "...my_host..."
    hostname_in_prefix          = true
    metrics = [
      {
        consumer_identifier  = "username"
        name                 = "shdict_usage"
        sample_rate          = 3.92
        service_identifier   = "service_name"
        stat_type            = "set"
        workspace_identifier = "workspace_name"
      }
    ]
    port   = 6303
    prefix = "...my_prefix..."
    queue = {
      concurrency_limit    = 0
      initial_retry_delay  = 75641.56
      max_batch_size       = 347740
      max_bytes            = 9
      max_coalescing_delay = 1430.25
      max_entries          = 139203
      max_retry_delay      = 226190.6
      max_retry_time       = 1.03
    }
    queue_size                   = 10
    retry_count                  = 8
    service_identifier_default   = "service_host"
    tag_style                    = "librato"
    udp_packet_size              = 61921
    use_tcp                      = false
    workspace_identifier_default = "workspace_name"
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
    "udp"
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