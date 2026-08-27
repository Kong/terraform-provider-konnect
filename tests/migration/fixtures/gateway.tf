# Fixture for scripts/workspace-migration-test.sh.
#
# Deliberately a pinned copy of tests/e2e/gateway.tf rather than a reference to
# it, for two reasons:
#
#   1. It must stay applicable by the OLD provider. If it tracked tests/e2e, a
#      new provider-only attribute added there would break phase 1 in a
#      confusing way. This file represents config a user wrote *before* the
#      upgrade, so it should not gain new syntax.
#   2. The control plane name below is distinct. Konnect scopes control plane
#      names per org, and tests/resources/testdata/TestGatewayControlPlane
#      already claims "Terraform Control Plane" - the acceptance job runs in
#      parallel with the migration job and would collide.
#
# Do not add a `workspace` argument here; the whole point is that it is absent.

# Create a new Control Plane
resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Migration Test Control Plane"
  description  = "Fixture for the workspace state-migration test"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
  auth_type    = "pinned_client_certs"

  proxy_urls = [
    {
      host     = "example.com",
      port     = 443,
      protocol = "https"
    }
  ]
}

# Configure a service and a route that we can use to test
resource "konnect_gateway_service" "httpbin" {
  name             = "HTTPBin"
  protocol         = "https"
  host             = "httpbin.org"
  port             = 443
  path             = "/"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_route" "hello" {
  methods = ["GET"]
  name    = "Anything"
  paths   = ["/anything"]
  headers = {
    Accept = ["application/json"]
  }

  strip_path = false

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

# Secure the service with a basic-auth plugin
resource "konnect_gateway_plugin_basic_auth" "basic_auth" {
  enabled          = true
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  service = {
    id = konnect_gateway_service.httpbin.id
  }
  config = {
    hide_credentials = false
  }
}

# Create a consumer and a basic auth credential for that consumer
resource "konnect_gateway_consumer" "alice" {
  username         = "alice"
  custom_id        = "alice"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}


# And a consumer group + membership
resource "konnect_gateway_consumer_group" "gold" {
  name             = "gold"
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_consumer_group_member" "ag" {
  consumer_id       = konnect_gateway_consumer.alice.id
  consumer_group_id = konnect_gateway_consumer_group.gold.id
  control_plane_id  = konnect_gateway_control_plane.tfdemo.id
}

# A plugin on every entity type
resource "konnect_gateway_plugin_rate_limiting" "global_rl" {
  enabled = true
  config = {
    minute = 100
    policy = "local"
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_plugin_rate_limiting" "service_rl" {
  enabled = true
  config = {
    minute = 1
    policy = "local"
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  service = {
    id = konnect_gateway_service.httpbin.id
  }
}

resource "konnect_gateway_plugin_rate_limiting" "route_rl" {
  enabled = true
  config = {
    minute = 2
    policy = "local"
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  route = {
    id = konnect_gateway_route.hello.id
  }
}

resource "konnect_gateway_plugin_rate_limiting" "consumer_rl" {
  enabled = true
  config = {
    minute = 3
    policy = "local"
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  consumer = {
    id = konnect_gateway_consumer.alice.id
  }
}

resource "konnect_gateway_plugin_rate_limiting" "consumer_group_rl" {
  enabled = true
  config = {
    minute = 4
    policy = "local"
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  consumer_group = {
    id = konnect_gateway_consumer_group.gold.id
  }
}

# A custom plugin schema
resource "konnect_gateway_custom_plugin_schema" "foo" {
  lua_schema = <<EOF
return {
  name = "myplugin",
  fields = {
    {
      config = {
        type = "record",
        fields = {
          { hello = { type = "string", required = true } },
        }
      }
    }
  }
}
EOF

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# This is a built in plugin, but we're using the custom_plugin resource
# to ensure that it works
resource "konnect_gateway_custom_plugin" "custom_basic_auth" {
  name             = "basic-auth"
  instance_name    = "custom-plugin-test"
  config           = {}
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Authentication types. We use the nested endpoint to ensure
# that consumer_id can be set

resource "konnect_gateway_acl" "my_acl" {
  group = "internal_users"

  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_basic_auth" "my_basicauth" {
  username = "alice-test"
  password = "demo"

  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_hmac_auth" "my_hmac" {
  username         = "alice"
  secret           = "secret1234"
  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_jwt" "my_jwt" {
  algorithm        = "HS256"
  secret           = "my_secret_value"
  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_key_auth" "my_keyauth" {
  key              = "abc123"
  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_mtls_auth" "my_mtlsauth" {
  subject_name     = "example.com"
  consumer_id      = konnect_gateway_consumer.alice.id
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Config store and vault
resource "konnect_gateway_config_store" "my_configstore" {
  name = "tf-config-store"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_vault" "my_vault" {
  name   = "konnect"
  prefix = "my-konnect-vault"
  config = jsonencode({
    config_store_id = konnect_gateway_config_store.my_configstore.id
  })
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
