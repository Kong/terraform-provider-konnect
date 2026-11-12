resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Credential Test Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_consumer" "alice" {
  username         = "alice"
  custom_id        = "alice"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_basic_auth" "my_basicauth" {
  username = "alice-test"
  password = "demo"

  consumer = {
    id = konnect_gateway_consumer.alice.id
  }

  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
