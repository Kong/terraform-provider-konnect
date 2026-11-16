resource "konnect_gateway_control_plane" "group" {
  name         = "Terraform Control Plane Group Membership"
  description  = "Control plane group for membership acceptance test"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE_GROUP"
}

resource "konnect_gateway_control_plane" "member" {
  name         = "Terraform Control Plane Group Member"
  description  = "Control plane member for membership acceptance test"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_control_plane_membership" "group" {
  control_plane_id = konnect_gateway_control_plane.group.id
  members = [{
    id = konnect_gateway_control_plane.member.id
  }]
}
