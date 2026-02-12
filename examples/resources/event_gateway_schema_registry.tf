resource "konnect_event_gateway_schema_registry" "sr" {
  gateway_id = konnect_event_gateway.demo.id

  confluent = {
    name = "sr-1"
    config = {
      endpoint    = "https://example-sr.local"
      schema_type = "json"
    }
  }
}
