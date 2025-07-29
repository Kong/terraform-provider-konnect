resource "konnect_gateway_control_plane" "plugincp" {
  name         = "Terraform Control Plane For Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_service" "my_service_nullable_for_plugin" {
  name             = "HTTPBin"
  protocol         = "https"
  host             = "httpbin.org"
  port             = 443
  path             = "/"
  control_plane_id = konnect_gateway_control_plane.plugincp.id
}

resource "konnect_gateway_route" "my_route_nullable_for_plugin" {
  methods = ["GET"]
  name    = "my-route-name"
  service = {
    id = konnect_gateway_service.my_service_nullable_for_plugin.id
  }
  control_plane_id = konnect_gateway_control_plane.plugincp.id
}

resource "konnect_gateway_consumer" "alice" {
  username         = "alice"
  control_plane_id = konnect_gateway_control_plane.plugincp.id
}

resource "konnect_gateway_plugin_rate_limiting" "my_nullable_rate_limiting" {
  enabled = true
  instance_name = "nullable-rl-plugin"
  config = {
    hour = 10000
    policy = "local"
  }
  service = {
    id = konnect_gateway_service.my_service_nullable_for_plugin.id
  }
  route = {
    id = konnect_gateway_route.my_route_nullable_for_plugin.id
  }

  consumer = {
    id = konnect_gateway_consumer.alice.id
  }

  // partials: to do - there is non-refresh plan diff which breaks tests.
  control_plane_id = konnect_gateway_control_plane.plugincp.id
}
