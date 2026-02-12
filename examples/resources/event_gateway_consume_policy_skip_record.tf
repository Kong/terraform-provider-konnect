resource "konnect_event_gateway_consume_policy_skip_record" "skip" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id
}
