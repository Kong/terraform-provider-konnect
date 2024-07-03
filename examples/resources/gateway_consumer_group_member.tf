resource "konnect_gateway_consumer_group_member" "ag" {
  consumer_id       = konnect_gateway_consumer.alice.id
  consumer_group_id = konnect_gateway_consumer_group.gold.id
  control_plane_id  = konnect_gateway_control_plane.tfdemo.id
}
