resource "konnect_event_gateway_consume_policy_skip_record" "my_eventgatewayconsumepolicyskiprecord" {
  condition   = "record.value.content.foo.bar == \"a-value\""
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "ed46f132-b990-42ef-9430-a541cbd0a3f7"
  virtual_cluster_id = "d2c285d9-943d-4169-81ee-6317de3cc511"
}