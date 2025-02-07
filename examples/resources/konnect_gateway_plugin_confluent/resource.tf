resource "konnect_gateway_plugin_confluent" "my_gatewaypluginconfluent" {
  config = {
    bootstrap_servers = [
      {
        host = "...my_host..."
        port = 53037
      }
    ]
    cluster_api_key                                    = "...my_cluster_api_key..."
    cluster_api_secret                                 = "...my_cluster_api_secret..."
    cluster_name                                       = "...my_cluster_name..."
    confluent_cloud_api_key                            = "...my_confluent_cloud_api_key..."
    confluent_cloud_api_secret                         = "...my_confluent_cloud_api_secret..."
    forward_body                                       = false
    forward_headers                                    = false
    forward_method                                     = false
    forward_uri                                        = true
    keepalive                                          = 3
    keepalive_enabled                                  = true
    producer_async                                     = false
    producer_async_buffering_limits_messages_in_memory = 9
    producer_async_flush_timeout                       = 1
    producer_request_acks                              = 0
    producer_request_limits_bytes_per_request          = 7
    producer_request_limits_messages_per_request       = 6
    producer_request_retries_backoff_timeout           = 4
    producer_request_retries_max_attempts              = 2
    producer_request_timeout                           = 1
    timeout                                            = 0
    topic                                              = "...my_topic..."
  }
  consumer = {
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