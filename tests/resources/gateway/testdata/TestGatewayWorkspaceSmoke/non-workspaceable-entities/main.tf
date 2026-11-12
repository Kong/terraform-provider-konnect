resource "konnect_gateway_control_plane" "smoke_nonworkspaceable_cp" {
  name         = "Terraform Control Plane For Non-Workspaceable Smoke Test"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

# Present to prove non-workspaceable entities are unaffected by, and don't
# require, a workspace even when one exists on the same control plane.
resource "konnect_gateway_workspace" "smoke" {
  name        = "tf-acceptance-smoke-workspace"
  description = "A workspace that coexists with non-workspaceable entities"

  control_plane_id = konnect_gateway_control_plane.smoke_nonworkspaceable_cp.id
}

# A representative sample of commonly used entities that are NOT
# workspaceable: they are always addressed directly through the control
# plane and have no "workspace" attribute at all.

resource "konnect_gateway_config_store" "smoke" {
  name = "smoke-config-store"

  control_plane_id = konnect_gateway_control_plane.smoke_nonworkspaceable_cp.id
}

resource "konnect_gateway_config_store_secret" "smoke" {
  key   = "smoke-key"
  value = "smoke-value"

  config_store_id  = konnect_gateway_config_store.smoke.id
  control_plane_id = konnect_gateway_control_plane.smoke_nonworkspaceable_cp.id
}

resource "konnect_gateway_ca_certificate" "smoke" {
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

  control_plane_id = konnect_gateway_control_plane.smoke_nonworkspaceable_cp.id
}

resource "konnect_gateway_custom_plugin_streaming" "smoke" {
  name = "setheader-streaming"

  handler = <<-EOT
    return {
      VERSION = "1.0.0",
      PRIORITY = 500,
      access = function(self, config)
        kong.service.request.set_header(config.name, config.value)
      end
    }
  EOT

  schema = <<-EOT
    return {
      name = "setheader-streaming",
      fields = {
        { protocols = require("kong.db.schema.typedefs").protocols_http },
        {
          config = {
            type = "record",
            fields = {
              { name = { description = "The name of the header to set.", type = "string", required = true, }, },
              { value = { description = "The value for the header.", type = "string", required = true, }, },
            },
          },
        },
      }
    }
  EOT

  control_plane_id = konnect_gateway_control_plane.smoke_nonworkspaceable_cp.id
}

resource "konnect_gateway_cloned_plugin" "smoke" {
  name = "smoke-custom-acl"
  ref  = "acl"

  control_plane_id = konnect_gateway_control_plane.smoke_nonworkspaceable_cp.id
}
