resource "konnect_gateway_plugin_ai_mcp_oauth2" "my_gatewaypluginaimcpoauth2" {
  config = {
    args = {
      key = "value"
    }
    authorization_servers = [
      "..."
    ]
    cache_introspection = false
    claim_to_header = [
      {
        claim  = "...my_claim..."
        header = "...my_header..."
      }
    ]
    client_alg    = "HS384"
    client_auth   = "none"
    client_id     = "...my_client_id..."
    client_jwk    = "...my_client_jwk..."
    client_secret = "...my_client_secret..."
    headers = {
      key = "value"
    }
    http_proxy                           = "...my_http_proxy..."
    http_proxy_authorization             = "...my_http_proxy_authorization..."
    http_version                         = 9.95
    https_proxy                          = "...my_https_proxy..."
    https_proxy_authorization            = "...my_https_proxy_authorization..."
    insecure_relaxed_audience_validation = false
    introspection_endpoint               = "...my_introspection_endpoint..."
    introspection_format                 = "base64"
    keepalive                            = false
    max_request_body_size                = 6
    metadata_endpoint                    = "...my_metadata_endpoint..."
    mtls_introspection_endpoint          = "...my_mtls_introspection_endpoint..."
    no_proxy                             = "...my_no_proxy..."
    resource                             = "...my_resource..."
    scopes_supported = [
      "..."
    ]
    ssl_verify                 = true
    timeout                    = 4.02
    tls_client_auth_cert       = "...my_tls_client_auth_cert..."
    tls_client_auth_key        = "...my_tls_client_auth_key..."
    tls_client_auth_ssl_verify = false
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 6
  enabled          = true
  id               = "...my_id..."
  instance_name    = "...my_instance_name..."
  ordering = {
    after = {
      access = [
        "..."
      ]
    }
    before = {
      access = [
        "..."
      ]
    }
  }
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
  protocols = [
    "grpc"
  ]
  route = {
    id = "...my_id..."
  }
  service = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
  updated_at = 4
}