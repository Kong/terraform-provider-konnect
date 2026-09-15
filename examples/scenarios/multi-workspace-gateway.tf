# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)

# Scopes Gateway core entities to a dedicated workspace instead of "default".
# `workspace` is Optional+Computed on every resource below, so it's safe to
# point it at a workspace created here and have each entity land there
# rather than in the control plane's default workspace.

# Create a new Control Plane
resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

# Create a dedicated workspace on that control plane
resource "konnect_gateway_workspace" "team_payments" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  name             = "team-payments"
  description      = "Workspace for the payments team"
}

# Configure a service and a route in the "team-payments" workspace
resource "konnect_gateway_service" "httpbin" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  name     = "HTTPBin"
  protocol = "https"
  host     = "httpbin.org"
  port     = 443
  path     = "/"
}

resource "konnect_gateway_route" "anything" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  methods    = ["GET"]
  name       = "Anything"
  paths      = ["/anything"]
  strip_path = false

  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

# Same workspace, but matched by an expression instead of paths/methods
resource "konnect_gateway_route_expression" "anything_expr" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  name       = "Anything Expression"
  expression = "http.path == \"/anything-expr\""
  protocols  = ["http", "https"]

  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

# Register a custom plugin's schema with the control plane (schemas are
# control-plane-wide, not workspace-scoped)
resource "konnect_gateway_custom_plugin_schema" "header_injector_schema" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id

  lua_schema = <<-EOT
    return {
      name = "header-injector",
      fields = {
        { protocols = require("kong.db.schema.typedefs").protocols_http },
        {
          config = {
            type = "record",
            fields = {
              { name  = { type = "string", required = true } },
              { value = { type = "string", required = true } },
            },
          },
        },
      },
    }
  EOT
}

# Create an instance of the custom plugin in the "team-payments" workspace,
# scoped to our service
resource "konnect_gateway_custom_plugin" "inject_header" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  name          = "header-injector"
  instance_name = "inject-custom-header"
  enabled       = true

  config = {
    name  = "x-custom-header"
    value = "my-custom-value"
  }

  service = {
    id = konnect_gateway_service.httpbin.id
  }

  depends_on = [konnect_gateway_custom_plugin_schema.header_injector_schema]
}

# Secure the service with key-auth, scoped to "team-payments"
resource "konnect_gateway_plugin_key_auth" "key_auth" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  enabled = true
  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

# Create a consumer in the same workspace, and a key-auth credential for it
resource "konnect_gateway_consumer" "alice" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  username  = "alice"
  custom_id = "alice"
}

resource "konnect_gateway_key_auth" "alice_key" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  consumer_id = konnect_gateway_consumer.alice.id
}

# Group the consumer via ACL, still scoped to "team-payments"
resource "konnect_gateway_acl" "alice_acl" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  consumer_id = konnect_gateway_consumer.alice.id
  group       = "gold"
}

# Enforce the ACL group above: only consumers in the "gold" group may call the service
resource "konnect_gateway_plugin_acl" "gold_only" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  enabled = true
  config = {
    allow = ["gold"]
  }

  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

# Cloned plugins are another way to get a custom plugin name: instead of
# defining a schema from scratch, clone an existing plugin (built-in or
# custom) under a new name. Like custom_plugin_schema, the clone itself is
# control-plane-wide, not workspace-scoped. Here we clone rate-limiting so
# "team-payments" can run its own distinctly-named instance of Kong's most
# widely used plugin.
resource "konnect_gateway_cloned_plugin" "team_payments_rate_limiting" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id

  name = "team-payments-rate-limiting"
  ref  = "rate-limiting"
}

# Instantiate the cloned plugin in the "team-payments" workspace, same as any
# other custom plugin
resource "konnect_gateway_custom_plugin" "rate_limit_route" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  name = "team-payments-rate-limiting"

  config = {
    minute = 100
    policy = "local"
  }

  route = {
    id = konnect_gateway_route.anything.id
  }

  depends_on = [konnect_gateway_cloned_plugin.team_payments_rate_limiting]
}
