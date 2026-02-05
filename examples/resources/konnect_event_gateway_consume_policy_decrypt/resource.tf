resource "konnect_event_gateway_consume_policy_decrypt" "my_eventgatewayconsumepolicydecrypt" {
  condition = "context.topic.name.endsWith(\"my_suffix\") && records.headers[\"x-flag\"] == \"a-value\""
  config = {
    failure_mode = "passthrough"
    key_sources = [
      {
        aws = {
          # ...
        }
        static = {
          # ...
        }
      }
    ]
    part_of_record = [
      "key"
    ]
  }
  description = "...my_description..."
  enabled     = false
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  virtual_cluster_id = "05c6c607-3c42-45e9-a9e8-3e6338120724"
}