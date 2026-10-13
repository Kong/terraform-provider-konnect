resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane For Gateway Workspace"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "my_gatewayworkspace" {
  name        = "tf-acceptance-workspace"
  description = "A test workspace for team 1"
  comment     = "A test workspace for team 1"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Created in the workspace above rather than the control plane's "default"
# workspace, so the workspace-scoped entity paths get exercised too.
resource "konnect_gateway_service" "httpbin" {
  name     = "HTTPBin"
  protocol = "https"
  host     = "httpbin.org"
  port     = 443
  path     = "/"

  workspace        = konnect_gateway_workspace.my_gatewayworkspace.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Everything below lives in the workspace above rather than "default". These
# exercise the workspace-scoped API paths, which differ from the control
# plane's default-workspace paths - and since both steps of this test apply
# the same config, the second step also proves they read back idempotently.

resource "konnect_gateway_route" "anything" {
  name      = "Anything"
  paths     = ["/anything"]
  protocols = ["http", "https"]

  service = {
    id = konnect_gateway_service.httpbin.id
  }

  workspace        = konnect_gateway_workspace.my_gatewayworkspace.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_consumer" "alice" {
  username  = "alice"
  custom_id = "alice"

  workspace        = konnect_gateway_workspace.my_gatewayworkspace.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Credentials are addressed under the consumer, so this covers the nested
# workspace-scoped path rather than a top-level collection.
resource "konnect_gateway_key_auth" "alice_key" {
  key         = "tf-acceptance-workspace-key"
  consumer_id = konnect_gateway_consumer.alice.id

  workspace        = konnect_gateway_workspace.my_gatewayworkspace.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# A plugin with no entity attached: scoped to the workspace itself.
resource "konnect_gateway_plugin_rate_limiting" "workspace_rl" {
  enabled = true
  config = {
    minute = 100
    policy = "local"
  }

  workspace        = konnect_gateway_workspace.my_gatewayworkspace.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# A plugin attached to an entity that itself lives in the workspace.
resource "konnect_gateway_plugin_cors" "service_cors" {
  enabled = true
  config = {
    origins = ["https://example.com"]
    methods = ["GET", "POST"]
    headers = ["Accept", "Content-Type"]
  }

  service = {
    id = konnect_gateway_service.httpbin.id
  }

  workspace        = konnect_gateway_workspace.my_gatewayworkspace.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
