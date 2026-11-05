# mask_fields requires an Event Gateway running runtime 1.3 or later.
# Set min_runtime_version = "1.3" on konnect_event_gateway.demo for this example.

# mask_fields is a child of a schema_validation policy and is attached to it
# through parent_policy_id.
resource "konnect_event_gateway_produce_policy_schema_validation" "produce_schema" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  name        = "test-produce-schema-validation-policy"
  description = "Test produce policy for schema validation"
  condition   = "context.topic.name == 'validated-topic'"
  enabled     = true

  config = {
    json = {
      key_validation_action   = "mark"
      value_validation_action = "reject"
    }
  }
}

resource "konnect_event_gateway_produce_policy_mask_fields" "mask_fields" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id
  parent_policy_id   = konnect_event_gateway_produce_policy_schema_validation.produce_schema.id

  name        = "tf-test-produce-mask-fields"
  description = "Mask personally identifiable fields before producing records"
  condition   = "record.value.content.foo.bar == 'a-value'"
  enabled     = true

  config = {
    failure_mode = "passthrough"
    mask_fields = [
      {
        paths = [
          {
            match = "user.card.number"
          }
        ]
        strategy = {
          replace = {
            phrase = "[redacted]"
          }
        }
      }
    ]
  }
}
