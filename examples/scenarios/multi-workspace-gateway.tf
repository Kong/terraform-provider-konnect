resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_workspace" "team_payments" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  name             = "team-payments"
  description      = "Workspace for the payments team"
}

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

resource "konnect_gateway_route_expression" "anything_expr" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  name       = "Anything-Expression"
  expression = "http.path == \"/anything-expr\""
  protocols  = ["http", "https"]

  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

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

resource "konnect_gateway_plugin_key_auth" "key_auth" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  enabled = true
  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

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

resource "konnect_gateway_acl" "alice_acl" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  consumer_id = konnect_gateway_consumer.alice.id
  group       = "gold"
}

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

resource "konnect_gateway_cloned_plugin" "team_payments_request_transformer" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id

  name = "team-payments-request-transformer"
  ref  = "request-transformer"
}

resource "konnect_gateway_custom_plugin" "request_transformer_route" {
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  workspace        = konnect_gateway_workspace.team_payments.name

  name = "team-payments-request-transformer"

  config = {
    add = {
      headers = ["x-team-payments: true"]
    }
  }

  route = {
    id = konnect_gateway_route.anything.id
  }

  depends_on = [konnect_gateway_cloned_plugin.team_payments_request_transformer]
}
