resource "konnect_event_gateway" "demo" {
  name                = "tf-event-gateway"
  min_runtime_version = "1.2"
}

resource "konnect_event_gateway_backend_cluster" "backend" {
  name              = "backend-1"
  bootstrap_servers = ["broker:9092"]
  gateway_id        = konnect_event_gateway.demo.id

  authentication = {
    anonymous = {}
  }
  tls = { enabled = false }
}

resource "konnect_event_gateway_virtual_cluster" "vc" {
  name       = "vc-1"
  dns_label  = "vcluster-1"
  acl_mode   = "passthrough"
  gateway_id = konnect_event_gateway.demo.id

  authentication = [
    {
      anonymous = {}
    }
  ]

  destination = { id = konnect_event_gateway_backend_cluster.backend.id }

  topic_aliases = [
    {
      alias    = "my-alias-topic"
      topic    = "my-backend-topic"
      conflict = "warn"
    }
  ]
}

# decrypt_fields uses key_sources with an empty static = {}, which resolves
# against static keys present in the gateway, so at least one must exist.
resource "konnect_event_gateway_static_key" "statickey" {
  name       = "tf-test-static-key"
  value      = "$${vault.env[\"MY_ENV_VAR\"]}"
  gateway_id = konnect_event_gateway.demo.id
}

# decrypt_fields is a child of a schema_validation policy and is attached to it
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

resource "konnect_event_gateway_consume_policy_decrypt_fields" "decrypt_fields" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id
  parent_policy_id   = konnect_event_gateway_consume_policy_schema_validation.consume_schema.id

  config = {
    failure_mode = "mark"
    key_sources = [
      {
        static = {}
      }
    ]
    decrypt_fields = {
      paths = [
        {
          match = "someObject.someArray[1].fieldName"
        }
      ]
    }
  }
}
