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

# ---------------------------------------------------------------------------
# Structural coverage.
#
# Everything above exercises flat strings, numbers and shallow objects. The
# upgrader round-trips state through map[string]any, so the shapes that can
# actually break are the awkward ones: multi-line strings, deeply nested
# objects, lists of numbers, and JSON embedded in a string attribute. The
# resources below exist to put those shapes in a real state file.
#
# All of them are supported by the baseline provider - nothing here may use
# post-migration syntax.
# ---------------------------------------------------------------------------

# PEM bodies are multi-line strings. If newline handling regressed anywhere in
# the JSON round-trip, this is where it surfaces as a diff.
resource "konnect_gateway_certificate" "migration_cert" {
  cert = <<EOF
-----BEGIN CERTIFICATE-----
MIIBxjCCAUygAwIBAgIUX9TaLbWF76yQc8IGR+YRbeiDlHkwCgYIKoZIzj0EAwIw
GjEYMBYGA1UEAwwPa29uZ19jbHVzdGVyaW5nMB4XDTI0MDMwMTE0MzkxNloXDTI3
MDMwMTE0MzkxNlowGjEYMBYGA1UEAwwPa29uZ19jbHVzdGVyaW5nMHYwEAYHKoZI
zj0CAQYFK4EEACIDYgAEcMndCotXzeZ9vGAMfDfZ7UxUuP5bcIrwwUOI8YlpMdvB
12HvjtS7O0/ONr3fBeCWagRuitPEqd4b3EJuD8kuFUMt+2A09N6KY1YDJWgKHei7
rzKgrefzVt11XgBiDsUBo1MwUTAdBgNVHQ4EFgQUIrdAC8p02h60GZW0Jlh2Vcg/
WeMwHwYDVR0jBBgwFoAUIrdAC8p02h60GZW0Jlh2Vcg/WeMwDwYDVR0TAQH/BAUw
AwEB/zAKBggqhkjOPQQDAgNoADBlAjBYb+yQf33sItlmsONLc41Agtx73FMEN7Lf
WA85OtlkMie1N1x0mj08pzS/Xc1VONwCMQDN9sBn3Kody0gse+EXYSuPPj1oo9jm
FB9/xrpz35YpDATvuyhH8xwSJ4xMuxQiduc=
-----END CERTIFICATE-----
EOF
  key  = <<EOF
-----BEGIN PRIVATE KEY-----
MIG2AgEAMBAGByqGSM49AgEGBSuBBAAiBIGeMIGbAgEBBDDLuRX+uzSbstvLWsQr
WwuGK4AdjLU/tN9A/fn03gxNvppKw++SBtnLyB+9YZ29YA+hZANiAARwyd0Ki1fN
5n28YAx8N9ntTFS4/ltwivDBQ4jxiWkx28HXYe+O1Ls7T842vd8F4JZqBG6K08Sp
3hvcQm4PyS4VQy37YDT03opjVgMlaAod6LuvMqCt5/NW3XVeAGIOxQE=
-----END PRIVATE KEY-----
EOF

  tags             = ["tf-migration", "structural"]
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# A required nested object referencing another resource's computed id.
resource "konnect_gateway_sni" "migration_sni" {
  name = "migration.example.com"
  certificate = {
    id = konnect_gateway_certificate.migration_cert.id
  }
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# The deepest nesting in the provider: healthchecks.active.healthy.http_statuses
# is an object -> object -> object -> list-of-number, four levels down.
resource "konnect_gateway_upstream" "migration_upstream" {
  name           = "migration-upstream.example.com"
  algorithm      = "round-robin"
  slots          = 10000
  hash_on        = "header"
  hash_on_header = "x-migration"

  healthchecks = {
    active = {
      concurrency              = 10
      http_path                = "/health"
      https_verify_certificate = true
      timeout                  = 1
      type                     = "http"
      healthy = {
        http_statuses = [200, 302]
        interval      = 5
        successes     = 2
      }
      unhealthy = {
        http_failures = 2
        http_statuses = [429, 500, 503]
        interval      = 5
        tcp_failures  = 2
        timeouts      = 2
      }
    }
    passive = {
      type = "http"
      healthy = {
        http_statuses = [200, 201, 302]
        successes     = 3
      }
      unhealthy = {
        http_failures = 3
        http_statuses = [500, 503]
        tcp_failures  = 3
        timeouts      = 3
      }
    }
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_target" "migration_target" {
  target      = "192.0.2.10:8080"
  weight      = 100
  upstream_id = konnect_gateway_upstream.migration_upstream.id
  tags        = ["tf-migration"]

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_gateway_key_set" "migration_key_set" {
  name             = "migration-key-set"
  tags             = ["tf-migration"]
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# `jwk` is JSON stored inside a string attribute - the quote- and
# brace-escaping case, distinct from the heredoc above.
resource "konnect_gateway_key" "migration_key" {
  kid  = "migration-kid-1"
  name = "migration-key"
  jwk = jsonencode({
    kty = "RSA"
    kid = "migration-kid-1"
    use = "sig"
    alg = "RS256"
    e   = "AQAB"
    n   = "tk5VuZUPrDvdDgS_Fouj2q1PXWyuYC0vEdKE6uZVrUrsoJlBAQZXNiMt3YtBvBUjXOEtRCXn7giRbHfqPxTFvNwKsiCIJZC7rsN6OCeQgC2hbiE77yAQtlWk_mWwvcVrg1Xo-ZcoxG592iJo9dTKRqpVmEWePNqIbUXJH4s3jjyPdkTqgfY9NZrBezilnAGgRlvDKJIye6iOqpw4KFqYaWcoiBTfgnTYhKGbhPrPRDFhspEfu6deTx069i84exMcExiaGRM-Bco7WfemnstkwGfi4u5jrVCQBg7r7TaBv37I4DY7uqOoWv1swobhM2Jk3vGRoLEZTiiIjPkTb2ZH6Q"
  })

  set = {
    id = konnect_gateway_key_set.migration_key_set.id
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Plugin config that is mostly lists of strings.
resource "konnect_gateway_plugin_cors" "migration_cors" {
  enabled = true
  config = {
    origins            = ["https://example.com", "https://*.example.org"]
    headers            = ["Accept", "Content-Type", "Authorization"]
    exposed_headers    = ["X-Migration-Test"]
    methods            = ["GET", "POST", "OPTIONS"]
    credentials        = true
    max_age            = 3600
    preflight_continue = false
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Nested objects whose every field is a list of strings.
resource "konnect_gateway_plugin_request_transformer" "migration_rt" {
  enabled = true
  config = {
    http_method = "POST"
    add = {
      headers     = ["x-migration:added"]
      querystring = ["migration:added"]
      body        = ["migration_body:added"]
    }
    append = {
      headers     = ["x-migration-append:appended"]
      querystring = ["appended:true"]
      body        = []
    }
    remove = {
      headers     = ["x-remove-me"]
      querystring = ["removeme"]
      body        = []
    }
    rename = {
      headers     = ["x-old:x-new"]
      querystring = []
      body        = []
    }
    replace = {
      headers     = ["x-migration:replaced"]
      querystring = []
      body        = []
      uri         = "/anything"
    }
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
  route = {
    id = konnect_gateway_route.hello.id
  }
}
