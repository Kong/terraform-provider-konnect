resource "konnect_gateway_plugin_statsd_advanced" "my_gatewaypluginstatsdadvanced" {
  config = {
    allow_status_codes = [
      "..."
    ]
    consumer_identifier_default = "username"
    host                        = "...my_host..."
    hostname_in_prefix          = false
    metrics = [
      {
        consumer_identifier  = "custom_id"
        name                 = "status_count_per_user"
        sample_rate          = 2.11
        service_identifier   = "service_id"
        stat_type            = "timer"
        workspace_identifier = "workspace_name"
      }
    ]
    port   = 36120
    prefix = "...my_prefix..."
    queue = {
      concurrency_limit    = 1
      initial_retry_delay  = 603656.13
      max_batch_size       = 216641
      max_bytes            = 3
      max_coalescing_delay = 1405.3
      max_entries          = 814148
      max_retry_delay      = 686342.11
      max_retry_time       = 8.95
    }
    service_identifier_default   = "service_host"
    udp_packet_size              = 19205.04
    use_tcp                      = false
    workspace_identifier_default = "workspace_id"
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
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
}