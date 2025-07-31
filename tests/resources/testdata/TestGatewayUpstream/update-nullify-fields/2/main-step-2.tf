resource "konnect_gateway_control_plane" "upstreamcp" {
  name         = "Terraform Control Plane For Upstream Nullify"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_upstream" "my_nullable_upstream" {
  name             = "demo-nullable-upstream"
  control_plane_id = konnect_gateway_control_plane.upstreamcp.id
}
