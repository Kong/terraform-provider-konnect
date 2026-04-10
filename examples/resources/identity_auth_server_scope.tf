resource "konnect_identity_auth_server_scope" "my_scope" {
  name                = "my-scope"
  description         = "My Scope"
  include_in_metadata = true
  enabled             = true

  auth_server_id = konnect_identity_auth_server.my_authserver.id
}
