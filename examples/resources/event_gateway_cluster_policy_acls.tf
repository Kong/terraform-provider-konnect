resource "konnect_event_gateway_virtual_cluster" "vc-acls" {
  name      = "tf-test-virtual-acls"
  dns_label = "tf-test-vc-acls"
  acl_mode  = "enforce_on_gateway"
  authentication = [
    {
      anonymous = {}
    }
  ]
  destination = {
    id = konnect_event_gateway_backend_cluster.backend.id
  }
  gateway_id         = konnect_event_gateway.demo.id
}

resource "konnect_event_gateway_cluster_policy_acls" "acls" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc-acls.id
  name        = "tf-test-cluster-policy-acls"
  description = "initial description"
  enabled     = false
  condition   = "context.topic.name.endsWith('my_suffix')"
  labels = {
    key = "value"
  }
  config = {
    rules = [
      {
        action = "deny"
        operations = [
          {
            name = "describe_configs"
          }
        ]
        resource_names = [
          {
            match = "...my_match..."
          }
        ]
        resource_type = "transactional_id"
      }
    ]
  }
}
