resource "konnect_event_gateway_produce_policy_modify_headers" "my_eventgatewayproducepolicymodifyheaders" {
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
  parent_policy_id   = "72c5778e-34a9-4e94-8979-28eb503453b5"
  virtual_cluster_id = "d146afe4-4af6-420a-9a5b-d37b93117501"
}