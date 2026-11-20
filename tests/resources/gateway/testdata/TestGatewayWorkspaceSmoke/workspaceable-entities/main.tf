resource "konnect_gateway_control_plane" "smoke_workspaceable_cp" {
  name         = "Terraform Control Plane For Workspace Smoke Test"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "smoke" {
  name        = "tf-acceptance-smoke-workspace"
  description = "A workspace used to smoke test workspaceable entities"

  control_plane_id = konnect_gateway_control_plane.smoke_workspaceable_cp.id
}

# A representative sample of commonly used entities that are workspaceable,
# all addressed through the non-default workspace above rather than the
# control plane's "default" one.

resource "konnect_gateway_service" "smoke" {
  name     = "HTTPBin"
  protocol = "https"
  host     = "httpbin.org"
  port     = 443
  path     = "/"

  workspace        = konnect_gateway_workspace.smoke.name
  control_plane_id = konnect_gateway_control_plane.smoke_workspaceable_cp.id
}

resource "konnect_gateway_route" "smoke" {
  name  = "smoke-route"
  paths = ["/anything"]
  service = {
    id = konnect_gateway_service.smoke.id
  }

  workspace        = konnect_gateway_workspace.smoke.name
  control_plane_id = konnect_gateway_control_plane.smoke_workspaceable_cp.id
}

resource "konnect_gateway_plugin_rate_limiting_advanced" "smoke" {
  enabled = true

  config = {
    limit       = [200]
    window_size = [1800]
    window_type = "fixed"
    namespace   = "smoke-namespace"
  }

  route = {
    id = konnect_gateway_route.smoke.id
  }

  workspace        = konnect_gateway_workspace.smoke.name
  control_plane_id = konnect_gateway_control_plane.smoke_workspaceable_cp.id
}
