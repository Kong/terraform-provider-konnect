resource "konnect_gateway_plugin_jwt_signer" "my_gatewaypluginjwtsigner" {
  config = {
    access_token_consumer_by = [
      "id"
    ]
    access_token_consumer_claim = [
      "..."
    ]
    access_token_introspection_authorization = "...my_access_token_introspection_authorization..."
    access_token_introspection_body_args     = "...my_access_token_introspection_body_args..."
    access_token_introspection_consumer_by = [
      "id"
    ]
    access_token_introspection_consumer_claim = [
      "..."
    ]
    access_token_introspection_endpoint = "...my_access_token_introspection_endpoint..."
    access_token_introspection_hint     = "...my_access_token_introspection_hint..."
    access_token_introspection_jwt_claim = [
      "..."
    ]
    access_token_introspection_leeway = 6.18
    access_token_introspection_scopes_claim = [
      "..."
    ]
    access_token_introspection_scopes_required = [
      "..."
    ]
    access_token_introspection_timeout       = 4.24
    access_token_issuer                      = "...my_access_token_issuer..."
    access_token_jwks_uri                    = "...my_access_token_jwks_uri..."
    access_token_jwks_uri_client_certificate = "...my_access_token_jwks_uri_client_certificate..."
    access_token_jwks_uri_client_password    = "...my_access_token_jwks_uri_client_password..."
    access_token_jwks_uri_client_username    = "...my_access_token_jwks_uri_client_username..."
    access_token_jwks_uri_rotate_period      = 0.18
    access_token_keyset                      = "...my_access_token_keyset..."
    access_token_keyset_client_certificate   = "...my_access_token_keyset_client_certificate..."
    access_token_keyset_client_password      = "...my_access_token_keyset_client_password..."
    access_token_keyset_client_username      = "...my_access_token_keyset_client_username..."
    access_token_keyset_rotate_period        = 4.53
    access_token_leeway                      = 0.51
    access_token_optional                    = false
    access_token_request_header              = "...my_access_token_request_header..."
    access_token_scopes_claim = [
      "..."
    ]
    access_token_scopes_required = [
      "..."
    ]
    access_token_signing_algorithm = "PS256"
    access_token_upstream_header   = "...my_access_token_upstream_header..."
    access_token_upstream_leeway   = 1.88
    add_access_token_claims = {
      key = jsonencode("value"),
    }
    add_channel_token_claims = {
      key = jsonencode("value"),
    }
    add_claims = {
      key = jsonencode("value"),
    }
    cache_access_token_introspection  = false
    cache_channel_token_introspection = true
    channel_token_consumer_by = [
      "username"
    ]
    channel_token_consumer_claim = [
      "..."
    ]
    channel_token_introspection_authorization = "...my_channel_token_introspection_authorization..."
    channel_token_introspection_body_args     = "...my_channel_token_introspection_body_args..."
    channel_token_introspection_consumer_by = [
      "id"
    ]
    channel_token_introspection_consumer_claim = [
      "..."
    ]
    channel_token_introspection_endpoint = "...my_channel_token_introspection_endpoint..."
    channel_token_introspection_hint     = "...my_channel_token_introspection_hint..."
    channel_token_introspection_jwt_claim = [
      "..."
    ]
    channel_token_introspection_leeway = 4.31
    channel_token_introspection_scopes_claim = [
      "..."
    ]
    channel_token_introspection_scopes_required = [
      "..."
    ]
    channel_token_introspection_timeout       = 6.9
    channel_token_issuer                      = "...my_channel_token_issuer..."
    channel_token_jwks_uri                    = "...my_channel_token_jwks_uri..."
    channel_token_jwks_uri_client_certificate = "...my_channel_token_jwks_uri_client_certificate..."
    channel_token_jwks_uri_client_password    = "...my_channel_token_jwks_uri_client_password..."
    channel_token_jwks_uri_client_username    = "...my_channel_token_jwks_uri_client_username..."
    channel_token_jwks_uri_rotate_period      = 9.27
    channel_token_keyset                      = "...my_channel_token_keyset..."
    channel_token_keyset_client_certificate   = "...my_channel_token_keyset_client_certificate..."
    channel_token_keyset_client_password      = "...my_channel_token_keyset_client_password..."
    channel_token_keyset_client_username      = "...my_channel_token_keyset_client_username..."
    channel_token_keyset_rotate_period        = 0.98
    channel_token_leeway                      = 4.86
    channel_token_optional                    = false
    channel_token_request_header              = "...my_channel_token_request_header..."
    channel_token_scopes_claim = [
      "..."
    ]
    channel_token_scopes_required = [
      "..."
    ]
    channel_token_signing_algorithm        = "PS384"
    channel_token_upstream_header          = "...my_channel_token_upstream_header..."
    channel_token_upstream_leeway          = 5.01
    enable_access_token_introspection      = false
    enable_channel_token_introspection     = true
    enable_hs_signatures                   = false
    enable_instrumentation                 = true
    original_access_token_upstream_header  = "...my_original_access_token_upstream_header..."
    original_channel_token_upstream_header = "...my_original_channel_token_upstream_header..."
    realm                                  = "...my_realm..."
    remove_access_token_claims = [
      "..."
    ]
    remove_channel_token_claims = [
      "..."
    ]
    set_access_token_claims = {
      key = jsonencode("value"),
    }
    set_channel_token_claims = {
      key = jsonencode("value"),
    }
    set_claims = {
      key = jsonencode("value"),
    }
    trust_access_token_introspection          = true
    trust_channel_token_introspection         = false
    verify_access_token_expiry                = true
    verify_access_token_introspection_expiry  = false
    verify_access_token_introspection_scopes  = false
    verify_access_token_scopes                = false
    verify_access_token_signature             = true
    verify_channel_token_expiry               = false
    verify_channel_token_introspection_expiry = false
    verify_channel_token_introspection_scopes = true
    verify_channel_token_scopes               = false
    verify_channel_token_signature            = false
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = false
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
  protocols = [
    "wss"
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
}