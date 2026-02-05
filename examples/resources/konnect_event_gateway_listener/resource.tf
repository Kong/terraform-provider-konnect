resource "konnect_event_gateway_listener" "my_eventgatewaylistener" {
  addresses = [
    "..."
  ]
  description = "...my_description..."
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name = "...my_name..."
  ports = [
    "..."
  ]
}