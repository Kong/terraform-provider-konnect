# Create a new Workspace inside the control plane defined in gateway.tf
resource "konnect_gateway_workspace" "team_1" {
  name             = "terraform-e2e-team-1"
  description      = "A test workspace for team 1"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Gateway entities are created in the "default" workspace unless the optional
# `workspace` argument names another one.
resource "konnect_gateway_service" "team_1_httpbin" {
  name     = "HTTPBin"
  protocol = "https"
  host     = "httpbin.org"
  port     = 443
  path     = "/"

  workspace        = konnect_gateway_workspace.team_1.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_route" "team_1_hello" {
  methods = ["GET"]
  name    = "Anything"
  paths   = ["/anything"]

  strip_path = false

  workspace        = konnect_gateway_workspace.team_1.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  service = {
    id = konnect_gateway_service.team_1_httpbin.id
  }
}

resource "konnect_gateway_consumer" "team_1_alice" {
  username  = "alice"
  custom_id = "alice"

  workspace        = konnect_gateway_workspace.team_1.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_plugin_rate_limiting" "team_1_rl" {
  enabled = true
  config = {
    minute = 5
    policy = "local"
  }

  workspace        = konnect_gateway_workspace.team_1.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  service = {
    id = konnect_gateway_service.team_1_httpbin.id
  }
}

resource "konnect_gateway_basic_auth" "team_1_basicauth" {
  username = "alice-team1"
  password = "demo"

  consumer_id      = konnect_gateway_consumer.team_1_alice.id
  workspace        = konnect_gateway_workspace.team_1.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_consumer_group" "team_1_gold" {
  name = "gold"

  workspace        = konnect_gateway_workspace.team_1.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_consumer_group_member" "team_1_ag" {
  consumer_id       = konnect_gateway_consumer.team_1_alice.id
  consumer_group_id = konnect_gateway_consumer_group.team_1_gold.id

  workspace        = konnect_gateway_workspace.team_1.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_upstream" "team_1_upstream" {
  name      = "team-1-upstream.example.com"
  algorithm = "round-robin"

  workspace        = konnect_gateway_workspace.team_1.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_target" "team_1_target" {
  target      = "192.0.2.10:8080"
  weight      = 100
  upstream_id = konnect_gateway_upstream.team_1_upstream.id

  workspace        = konnect_gateway_workspace.team_1.name
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
