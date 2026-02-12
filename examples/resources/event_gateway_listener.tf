resource "konnect_event_gateway_listener" "listener" {

  gateway_id = konnect_event_gateway.demo.id
  ports      = [9092]
  addresses  = ["0.0.0.0"]
  name       = "listener-1"
}
