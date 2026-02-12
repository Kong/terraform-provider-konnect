resource "konnect_event_gateway_consume_policy_schema_validation" "consume_schema" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  config = { type = "json" }
}
