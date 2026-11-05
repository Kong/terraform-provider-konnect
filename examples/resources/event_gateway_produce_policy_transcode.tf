# transcode requires an Event Gateway running runtime 1.3 or later.
# Set min_runtime_version = "1.3" on konnect_event_gateway.demo for this example.
resource "konnect_event_gateway_schema_registry" "sr" {
  gateway_id = konnect_event_gateway.demo.id

  confluent = {
    name        = "sr-1"
    description = "Schema registry used by the produce transcode policy"
    config = {
      endpoint    = "https://example-sr.local"
      schema_type = "avro"
    }
  }
}

# transcode is a child of a schema_validation policy and is attached to it
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

resource "konnect_event_gateway_produce_policy_transcode" "transcode" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id
  parent_policy_id   = konnect_event_gateway_produce_policy_schema_validation.produce_schema.id

  name        = "tf-test-produce-transcode"
  description = "Transcode produced records to Avro"
  condition   = "record.value.content.foo.bar == 'a-value'"
  enabled     = true

  config = {
    avro = {
      failure_mode     = "mark"
      on_unknown_field = "drop"
      schema_source = {
        reference = {
          schema_ref_destination = {
            confluent_format = {}
          }
          schema_registry = {
            id = konnect_event_gateway_schema_registry.sr.id
          }
          subject = "context.topic.name + '-value'"
          version = "'latest'"
        }
      }
    }
  }

  depends_on = [
    konnect_event_gateway_schema_registry.sr
  ]
}
