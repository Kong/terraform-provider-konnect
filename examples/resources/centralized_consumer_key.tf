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
