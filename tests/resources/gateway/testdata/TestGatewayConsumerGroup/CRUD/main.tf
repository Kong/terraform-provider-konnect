resource "konnect_gateway_control_plane" "gwconsumergroupcp" {
  name         = "Terraform Control Plane For Gateway Consumer Group"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_consumer" "my_consumer" {
  username         = "tfcgtestconsumer"
  control_plane_id = konnect_gateway_control_plane.gwconsumergroupcp.id
}

resource "konnect_gateway_consumer_group" "my_gatewayconsumergroup" {
  control_plane_id = konnect_gateway_control_plane.gwconsumergroupcp.id
  name             = "TFAcceptanceGWConsumerGroupName"
  tags = ["tag1", "tag2"]
}

resource "konnect_gateway_consumer_group_member" "my_gatewayconsumergroupmember" {
  consumer_group_id = konnect_gateway_consumer_group.my_gatewayconsumergroup.id
  consumer_id       = konnect_gateway_consumer.my_consumer.id
  control_plane_id  = konnect_gateway_control_plane.gwconsumergroupcp.id
}