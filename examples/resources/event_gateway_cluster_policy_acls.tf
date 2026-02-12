resource "konnect_event_gateway_cluster_policy_acls" "acls" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  config = {
    rules = [
      {
        action = "deny"
        resource_type = "topic"
        operations = [
          {
            name = "describe"
          }
        ]
        resource_names = [
          {
            match = "topic_*"
          }
        ]
      }
    ]
  }
}
