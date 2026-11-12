resource "konnect_gateway_control_plane" "centralized_consumer_cp" {
  name         = "Centralized Consumers Control Plane for null support"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_realm" "my_realm" {
  name                     = "nullify-consumer-id"
  ttl                      = 60
  negative_ttl             = 10
  allow_all_control_planes = false
  allowed_control_planes   = [konnect_gateway_control_plane.centralized_consumer_cp.id]
  force_destroy = "true"
}

resource "konnect_centralized_consumer" "my_centralizedconsumer" {
  username  = "alice"
  custom_id = "alice-id"
  realm_id  = konnect_realm.my_realm.id
}