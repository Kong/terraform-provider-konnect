resource "konnect_gateway_plugin_kafka_upstream" "my_gatewaypluginkafkaupstream" {
  config = {
    authentication = {
      mechanism = "SCRAM-SHA-512"
      password  = "...my_password..."
      strategy  = "sasl"
      tokenauth = false
      user      = "...my_user..."
    }
    bootstrap_servers = [
      {
        host = "...my_host..."
        port = 3969
      }
    ]
    cluster_name                                       = "...my_cluster_name..."
    forward_body                                       = false
    forward_headers                                    = true
    forward_method                                     = true
    forward_uri                                        = false
    keepalive                                          = 3
    keepalive_enabled                                  = true
    producer_async                                     = true
    producer_async_buffering_limits_messages_in_memory = 5
    producer_async_flush_timeout                       = 5
    producer_request_acks                              = 1
    producer_request_limits_bytes_per_request          = 1
    producer_request_limits_messages_per_request       = 9
    producer_request_retries_backoff_timeout           = 4
    producer_request_retries_max_attempts              = 5
    producer_request_timeout                           = 10
    security = {
      certificate_id = "...my_certificate_id..."
      ssl            = false
    }
    timeout = 0
    topic   = "...my_topic..."
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