resource "konnect_event_gateway_listener_policy_forward_to_virtual_cluster" "forward" {
  gateway_id  = konnect_event_gateway.demo.id
  listener_id = konnect_event_gateway_listener.listener.id

  config = { sni = { sni_suffix = ".example.com" } }
}
