resource "konnect_gateway_control_plane" "GWConsumerKeyAuthCP" {
  name         = "Terraform Control Plane For Vault"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_consumer" "KeyAuthConsumer" {
  username         = "nullableconsumer"
  custom_id        = "nullableConsumerID"
  control_plane_id = konnect_gateway_control_plane.GWConsumerKeyAuthCP.id
}

resource "konnect_gateway_key_auth" "my_keyauth" {
  key = "my-dummy-key"
  consumer_id      = konnect_gateway_consumer.KeyAuthConsumer.id
  control_plane_id = konnect_gateway_control_plane.GWConsumerKeyAuthCP.id
}