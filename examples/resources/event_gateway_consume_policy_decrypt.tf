resource "konnect_event_gateway_consume_policy_decrypt" "decrypt" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  config = {
    failure_mode   = "passthrough"
    part_of_record = ["key"]
    key_sources    = [
      {
        static =
          {
            key =
            {
              id = konnect_event_gateway_static_key.statickey.id
            }
          }
      }
    ]
  }
}
