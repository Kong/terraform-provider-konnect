# transcode requires an Event Gateway running runtime 1.3 or later.
# Set min_runtime_version = "1.3" on konnect_event_gateway.demo for this example.

# transcode is a child of a schema_validation policy and is attached to it
# through parent_policy_id.
resource "konnect_event_gateway_consume_policy_schema_validation" "consume_schema" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  name        = "test-consume-schema-validation-policy"
  description = "Test consume policy for schema validation"
  condition   = "context.topic.name == 'validated-topic'"
  enabled     = true

  config = {
    json = {
      key_validation_action   = "mark"
      value_validation_action = "skip"
    }
  }
}

resource "konnect_event_gateway_consume_policy_transcode" "transcode" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id
  parent_policy_id   = konnect_event_gateway_consume_policy_schema_validation.consume_schema.id

  name        = "tf-test-consume-transcode"
  description = "Transcode consumed records to JSON"
  condition   = "record.value.content.foo.bar == 'a-value'"
  enabled     = true

  config = {
    json = {
      failure_mode     = "mark"
      on_unknown_field = "drop"
      schema_source = {
        inline = {
          inline = "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"},\"user\":{\"type\":\"object\"}}}"
        }
      }
    }
  }
}
