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

resource "konnect_event_gateway_schema_registry" "sr" {
  gateway_id = konnect_event_gateway.demo.id

  confluent = {
    name        = "sr-1"
    description = "Schema registry for the encrypt_fields parent policy"
    config = {
      endpoint    = "https://example-sr.local"
      schema_type = "json"
    }
  }
}

resource "konnect_event_gateway_static_key" "statickey" {
  name       = "tf-test-static-key"
  value      = "$${vault.env[\"MY_ENV_VAR\"]}"
  gateway_id = konnect_event_gateway.demo.id
}

# encrypt_fields is a child of a schema_validation policy and is attached to it
# through parent_policy_id.
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

resource "konnect_event_gateway_produce_policy_encrypt_fields" "encrypt_fields" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id
  parent_policy_id   = konnect_event_gateway_produce_policy_schema_validation.produce_schema.id

  config = {
    failure_mode = "mark"
    encrypt_fields = [
      {
        encryption_key = {
          static = {
            key = {
              id = konnect_event_gateway_static_key.statickey.id
            }
          }
        }
        paths = [
          {
            match = "someObject.someArray[1].fieldName"
          }
        ]
      }
    ]
  }
}
