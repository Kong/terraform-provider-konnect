resource "konnect_gateway_control_plane" "expr_route_cp" {
  name         = "Terraform Control Plane For Expr Route"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_route_expression" "my_nullify_expression_route" {
  expression = "http.path ^= \"/anything\""
  strip_path = false

  control_plane_id = konnect_gateway_control_plane.expr_route_cp.id
}
