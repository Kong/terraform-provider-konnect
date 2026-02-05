resource "konnect_event_gateway_listener_policy_forward_to_virtual_cluster" "my_eventgatewaylistenerpolicyforwardtovirtualcluster" {
  config = {
    sni = {
      advertised_port = 61579
      broker_host_format = {
        type = "per_cluster_suffix"
      }
      sni_suffix = ".example.com"
    }
  }
  description = "...my_description..."
  enabled     = false
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  listener_id = "bdaf2651-42bc-48ec-b29f-f4890f7f07fc"
  name        = "...my_name..."
}