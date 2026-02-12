resource "konnect_event_gateway_consume_policy_modify_headers" "modify_consume" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  config = {
    actions = [
      {
        set = {
          key = "x-added", value = "v"
        }
      }
    ]
  }
}
