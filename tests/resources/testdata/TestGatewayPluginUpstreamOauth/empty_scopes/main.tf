resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Upstream OAuth Test CP"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_upstream_oauth" "test" {
  enabled = true

  config = {
    oauth = {
      token_endpoint = "http://test.test"
      grant_type = "client_credentials"
      client_id = "CLIENT_CREDENTIALS_GRANT_POST_AUTH_CLIENT_ID"
      client_secret = "CLIENT_CREDENTIALS_GRANT_POST_AUTH_CLIENT_SECRET"
    }
    behavior = {
      upstream_access_token_header_name = "X-Custom-Auth"
    }
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}