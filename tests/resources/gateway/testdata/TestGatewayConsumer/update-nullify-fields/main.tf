resource "konnect_gateway_control_plane" "gwconsumercp" {
  name         = "Terraform Control Plane For Gateway Consumer"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_consumer" "nullableconsumer" {
  username         = "nullableconsumer"
  custom_id        = "nullableConsumerID"
  control_plane_id = konnect_gateway_control_plane.gwconsumercp.id
}
