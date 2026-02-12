resource "konnect_event_gateway_backend_cluster" "backend" {
  name              = "backend-1"
  bootstrap_servers = ["broker:9092"]
  gateway_id        = konnect_event_gateway.demo.id

  authentication = {
    anonymous = {}
  }
  tls = { enabled = false }
}
