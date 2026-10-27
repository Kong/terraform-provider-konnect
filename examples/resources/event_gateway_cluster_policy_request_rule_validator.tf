# Request rule validation requires an Event Gateway running runtime 1.3 or later.
# Set min_runtime_version = "1.3" on konnect_event_gateway.demo for this example.
resource "konnect_event_gateway_cluster_policy_request_rule_validator" "request_rules" {
  gateway_id         = konnect_event_gateway.demo.id
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.vc.id

  name        = "tf-test-request-rule-validator"
  description = "Enforce the topic naming convention and consumer group ownership"
  enabled     = true
  condition   = "context.auth.principal.name == 'team-a-service'"
  labels = {
    env = "test"
  }

  config = {
    requests = [
      {
        create_topics = {
          rules = [
            {
              action      = "reject"
              description = "Enforce the topic naming convention and a partition limit."
              rule        = "context.topic.name.startsWith('team-a.') && context.topic.num_partitions <= 50"
            }
          ]
        }
      },
      {
        consumer_group_heartbeat = {
          rules = [
            {
              action      = "passthrough"
              description = "Log heartbeats sent by consumer groups outside the team-a prefix."
              rule        = "context.group.id.startsWith('team-a.')"
            }
          ]
        }
      }
    ]
  }
}
