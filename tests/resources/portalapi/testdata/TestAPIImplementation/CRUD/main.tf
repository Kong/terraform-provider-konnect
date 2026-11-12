resource "konnect_gateway_control_plane" "api_impl_cp" {
  name         = "API Implementation Test CP"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_control_plane" "my_cp_for_api_cp_link" {
  name         = "API Implementation linked CP"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

// Required to link API Implementation to Control Plane
resource "konnect_gateway_plugin_ace" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin-for-api-impl-cp-link"

  config = {
    match_policy = "required"
  }

  control_plane_id = konnect_gateway_control_plane.my_cp_for_api_cp_link.id
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

resource "konnect_api" "my_api_cp_linked" {
  description = "My API Description"
  name         = "MyCPLinkedAPI"
  slug         = "my-api-v1-cp-linked-implementation"
  version      = "v1"
}

resource "konnect_api_implementation" "my_apiimplementation" {
  api_id = konnect_api.my_api.id
  service_reference = {
    service = {
      control_plane_id = konnect_gateway_control_plane.api_impl_cp.id
      id               = konnect_gateway_service.httpbin.id
    }
  }
}

resource "konnect_api_implementation" "my_apiimplementation_cp_linked" {
  depends_on = [konnect_gateway_plugin_ace.my_plugin]
  api_id = konnect_api.my_api_cp_linked.id
  control_plane_reference = {
    control_plane = {
      id               = konnect_gateway_control_plane.my_cp_for_api_cp_link.id
    }
  }
}