resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Custom Plugin Test"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_service" "httpbin" {
  name             = "HTTPBin"
  protocol         = "https"
  host             = "httpbin.org"
  port             = 443
  path             = "/"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_custom_plugin" "custom_basic_auth" {
  name             = "basic-auth"
  instance_name    = "custom-plugin-test"
  config           = {}
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_custom_plugin" "custom_basic_auth_nested" {
  name             = "basic-auth"
  instance_name    = "custom-nested-plugin-test"
  config           = {}
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

resource "konnect_gateway_custom_plugin" "custom_no_instance_name" {
  name             = "key-auth"
  config           = {}
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
