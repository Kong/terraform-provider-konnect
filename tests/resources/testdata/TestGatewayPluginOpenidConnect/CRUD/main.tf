resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "OIDC CRUD Test CP"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_openid_connect" "test" {
  enabled = true

  config = {
    issuer = "https://example.com"
    consumer_claims = [
      ["sub"]
    ]
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
