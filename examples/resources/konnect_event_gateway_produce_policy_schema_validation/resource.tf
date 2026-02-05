resource "konnect_event_gateway_produce_policy_schema_validation" "my_eventgatewayproducepolicyschemavalidation" {
  condition = "context.topic.name.endsWith(\"my_suffix\") && records.headers[\"x-flag\"] == \"a-value\""
  config = {
    confluent_schema_registry = {
      key_validation_action = "reject"
      schema_registry = {
        id = "95ce4f52-159c-43dd-a6dd-9bb6e8e07446"
      }
      value_validation_action = "reject"
    }
    json = {
      key_validation_action = "mark"
      schema_registry = {
        id = "74577697-03b2-4d40-bfe2-929c891c4254"
      }
      value_validation_action = "reject"
    }
  }
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  virtual_cluster_id = "a3f4c612-4025-4392-861f-faa39b63e12d"
}