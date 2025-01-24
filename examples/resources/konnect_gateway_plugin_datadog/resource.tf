resource "konnect_gateway_plugin_datadog" "my_gatewayplugindatadog" {
  config = {
    consumer_tag  = "...my_consumer_tag..."
    flush_timeout = 5.18
    host          = "...my_host..."
    metrics = [
      {
        consumer_identifier = "username"
        name                = "request_size"
        sample_rate         = 0.12
        stat_type           = "timer"
        tags = [
          "..."
        ]
      }
    ]
    port   = 19191
    prefix = "...my_prefix..."
    queue = {
      concurrency_limit    = 1
      initial_retry_delay  = 632515.49
      max_batch_size       = 689907
      max_bytes            = 2
      max_coalescing_delay = 645.74
      max_entries          = 470545
      max_retry_delay      = 272290.78
      max_retry_time       = 9.63
    }
    queue_size       = 10
    retry_count      = 2
    service_name_tag = "...my_service_name_tag..."
    status_tag       = "...my_status_tag..."
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
    "ws"
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