resource "konnect_gateway_route_expression" "expression_route" {
  name       = "AnythingExpression"
  expression = "http.path ^= \"/anything\""
  strip_path = false

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  service = {
    id = konnect_gateway_service.httpbin.id
  }
}
