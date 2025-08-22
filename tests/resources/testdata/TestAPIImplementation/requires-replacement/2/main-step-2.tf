resource "konnect_gateway_control_plane" "api_impl_cp" {
  name         = "API Implementation Test CP"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_service" "httpbin" {
  protocol         = "https"
  host             = "httpbin.org"
  port             = 443
  control_plane_id = konnect_gateway_control_plane.api_impl_cp.id
}

resource "konnect_api" "my_api" {
  description = "My API Description"
  name         = "MyAPI"
  slug         = "my-api-v1-implementation"
  version      = "v1"
}

resource "konnect_api_implementation" "my_apiimplementation" {
  api_id = konnect_api.my_api.id
}