resource "konnect_centralized_consumer" "my_centralizedconsumer" {
  username  = "alice"
  custom_id = "u1234"
  realm_id  = konnect_realm.my_realm.id
}
