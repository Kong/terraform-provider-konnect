resource "konnect_event_gateway_produce_policy_encrypt" "encrypt" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  config = {
    encryption_key = { static = { key = { id = konnect_event_gateway_static_key.statickey.id } } }
    failure_mode   = "error"
    part_of_record = ["key"]
  }
}
