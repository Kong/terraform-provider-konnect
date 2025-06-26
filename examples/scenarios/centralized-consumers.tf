# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)

resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Centralized Consumers Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_realm" "my_realm" {
  name                     = "Production"
  ttl                      = 60
  negative_ttl             = 10
  allow_all_control_planes = false
  allowed_control_planes   = [konnect_gateway_control_plane.tfdemo.id]
  consumer_groups = [
    "prod-user"
  ]
  force_destroy = "true"
}

resource "konnect_centralized_consumer" "my_centralizedconsumer" {
  username  = "alice"
  custom_id = "u1234"
  realm_id  = konnect_realm.my_realm.id
}

resource "konnect_centralized_consumer_key" "my_centralizedconsumerkey" {
  type        = "new"
  consumer_id = konnect_centralized_consumer.my_centralizedconsumer.id
  realm_id    = konnect_realm.my_realm.id
}

resource "konnect_centralized_consumer_key" "my_centralizedconsumerkey_legacy" {
  type        = "legacy"
  secret      = "this-is-an-old-key-that-is-imported-this-is-not-recommended"
  consumer_id = konnect_centralized_consumer.my_centralizedconsumer.id
  realm_id    = konnect_realm.my_realm.id
}
