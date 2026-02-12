resource "konnect_event_gateway_consume_policy_schema_validation" "consume_schema" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  name        = "test-consume-schema-validation-policy"
  description = "Test consume policy for schema validation"
  condition   = "context.topic.name == 'validated-topic'"
  enabled     = true

  config = {
    type                     = "json"
    key_validation_action    = "mark"
    value_validation_action  = "skip"
  }
}
