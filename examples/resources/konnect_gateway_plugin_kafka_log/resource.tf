resource "konnect_gateway_plugin_kafka_log" "my_gatewaypluginkafkalog" {
  config = {
    authentication = {
      mechanism = "SCRAM-SHA-256"
      password  = "...my_password..."
      strategy  = "sasl"
      tokenauth = false
      user      = "...my_user..."
    }
    bootstrap_servers = [
      {
        host = "...my_host..."
        port = 7302
      }
    ]
    cluster_name = "...my_cluster_name..."
    custom_fields_by_lua = {
      key = jsonencode("value"),
    }
    keepalive                                          = 2
    keepalive_enabled                                  = true
    producer_async                                     = false
    producer_async_buffering_limits_messages_in_memory = 2
    producer_async_flush_timeout                       = 6
    producer_request_acks                              = 0
    producer_request_limits_bytes_per_request          = 5
    producer_request_limits_messages_per_request       = 10
    producer_request_retries_backoff_timeout           = 1
    producer_request_retries_max_attempts              = 4
    producer_request_timeout                           = 3
    security = {
      certificate_id = "...my_certificate_id..."
      ssl            = false
    }
    timeout = 2
    topic   = "...my_topic..."
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