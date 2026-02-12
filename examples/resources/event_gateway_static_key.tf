resource "konnect_event_gateway_static_key" "statickey" {
  name       = "static-key-1"
  value      = "REPLACE_ME"
  gateway_id = konnect_event_gateway.demo.id
}
