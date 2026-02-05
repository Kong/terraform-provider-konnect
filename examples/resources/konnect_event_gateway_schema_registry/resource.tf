resource "konnect_event_gateway_schema_registry" "my_eventgatewayschemaregistry" {
  confluent = {
    config = {
      authentication = {
        basic = {
          password = "$${vault.env['MY_ENV_VAR']}"
          username = "...my_username..."
        }
      }
      endpoint        = "https://key-hovercraft.com"
      schema_type     = "avro"
      timeout_seconds = 8
    }
    description = "...my_description..."
    labels = {
      key = "value"
    }
    name = "...my_name..."
  }
  gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
}