resource "konnect_event_gateway_produce_policy_schema_validation" "produce_schema" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  name        = "tf-test-produce-schema-validation"
  description = "initial schema validation policy"
  enabled     = true

  config = {
    confluent_schema_registry = {
      key_validation_action   = "reject"
      value_validation_action = "reject"
      schema_registry = {
        id = konnect_event_gateway_schema_registry.sr.id
      }
    }
  }
}
