resource "konnect_event_gateway_listener_policy_forward_to_virtual_cluster" "tf_test_egw_forward_policy" {
  name        = "tf_test_egw_forward_policy"
  description = "Test forward to virtual cluster policy"
  enabled     = true

  config = {
    port_mapping = {
      advertised_host = "broker.example.com"
      bootstrap_port  = "at_start"
      min_broker_id   = 0
      destination = {
        id = konnect_event_gateway_virtual_cluster.vc.id
      }
    }
  }
  gateway_id  = konnect_event_gateway.demo.id
  listener_id = konnect_event_gateway_listener.listener.id
}
