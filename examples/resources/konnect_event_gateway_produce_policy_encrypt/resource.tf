resource "konnect_event_gateway_produce_policy_encrypt" "my_eventgatewayproducepolicyencrypt" {
  condition = "context.topic.name.endsWith(\"my_suffix\") && records.headers[\"x-flag\"] == \"a-value\""
  config = {
    encryption_key = {
      static = {
        key = {
          reference_by_id = {
            id = "9d9dcdc8-beb0-45dc-8f7e-521cf4b6c0c7"
          }
        }
      }
    }
    failure_mode = "error"
    part_of_record = [
      "key"
    ]
  }
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  virtual_cluster_id = "6ea3798e-38ca-4c28-a68e-1a577e478f2c"
}