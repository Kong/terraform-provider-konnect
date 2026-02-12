resource "konnect_event_gateway_virtual_cluster" "vc" {
  name       = "vc-1"
  dns_label  = "vcluster-1"
  acl_mode   = "passthrough"
  gateway_id = konnect_event_gateway.demo.id

  authentication = [
    {
      anonymous = { }
    }
  ]

  destination = { id = konnect_event_gateway_backend_cluster.backend.id }
}
