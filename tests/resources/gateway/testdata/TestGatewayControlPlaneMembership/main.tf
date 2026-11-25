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

resource "konnect_gateway_control_plane" "member2" {
  name         = "Terraform Control Plane Group Member 2"
  description  = "Another control plane member for membership acceptance test"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_control_plane_membership" "group_member" {
  id = konnect_gateway_control_plane.group.id
  members = [{
    id = konnect_gateway_control_plane.member.id
  }]
}

resource "konnect_gateway_control_plane_membership" "group_member2" {
  id = konnect_gateway_control_plane.group.id
  members = [{
    id = konnect_gateway_control_plane.member2.id
  }]
}
