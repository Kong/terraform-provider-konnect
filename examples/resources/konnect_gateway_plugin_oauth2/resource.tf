resource "konnect_gateway_plugin_oauth2" "my_gatewaypluginoauth2" {
  config = {
    accept_http_if_already_terminated = true
    anonymous                         = "...my_anonymous..."
    auth_header_name                  = "...my_auth_header_name..."
    enable_authorization_code         = true
    enable_client_credentials         = false
    enable_implicit_grant             = true
    enable_password_grant             = true
    global_credentials                = true
    hide_credentials                  = false
    mandatory_scope                   = false
    persistent_refresh_token          = true
    pkce                              = "strict"
    provision_key                     = "...my_provision_key..."
    realm                             = "...my_realm..."
    refresh_token_ttl                 = 43385286.36
    reuse_refresh_token               = true
    scopes = [
      "..."
    ]
    token_expiration = 5.01
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
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