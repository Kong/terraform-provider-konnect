resource "konnect_event_gateway" "my_eventgateway" {
  description = "...my_description..."
  labels = {
    key = "value"
  }
  min_runtime_version = "1.1"
  name                = "...my_name..."
}