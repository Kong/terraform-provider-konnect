resource "konnect_gateway_jwt" "my_jwt" {
  algorithm        = "HS256"
  secret           = "my_secret_value"
  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
