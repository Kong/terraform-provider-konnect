resource "konnect_event_gateway_consume_policy_modify_headers" "my_eventgatewayconsumepolicymodifyheaders" {
  condition = "record.value.content.foo.bar == \"a-value\""
  config = {
    actions = [
      {
        remove = {
          key = "...my_key..."
        }
      }
    ]
  }
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "45b2f9d6-c646-4faa-8c5f-7d1ebf8687c1"
  virtual_cluster_id = "58221d12-f9c8-4032-9ae2-54155e337f04"
}