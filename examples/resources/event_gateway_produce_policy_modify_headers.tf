resource "konnect_event_gateway_produce_policy_modify_headers" "modify_produce" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  config = {
    actions = [{ remove = { key = "x-remove" } }]
  }
}
