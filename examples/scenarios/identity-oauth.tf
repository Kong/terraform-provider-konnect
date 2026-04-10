resource "konnect_identity_auth_server" "my_authserver" {
  name     = "my-auth-server"
  audience = "local-demo"
}

resource "konnect_identity_auth_server_scope" "my_scope" {
  name                = "my-scope"
  description         = "My Scope"
  include_in_metadata = true
  enabled             = true

  auth_server_id = konnect_identity_auth_server.my_authserver.id
}

resource "konnect_identity_auth_server_claim" "my_authserverclaims" {
  enabled               = true
  include_in_all_scopes = false
  include_in_scopes = [
    konnect_identity_auth_server_scope.my_scope.id
  ]
  include_in_token = true
  name             = "my-claim"
  value            = "some-value"

  auth_server_id = konnect_identity_auth_server.my_authserver.id
}

resource "konnect_identity_auth_server_client" "my_client" {
  name             = "my-client"
  allow_all_scopes = true
  grant_types = [
    "client_credentials",
    "implicit"
  ]

  response_types = [
    "id_token",
    "token",
    "code"
  ]

  auth_server_id = konnect_identity_auth_server.my_authserver.id
}

output "issuer" {
  value = konnect_identity_auth_server.my_authserver.issuer
}

output "client_id" {
  value = konnect_identity_auth_server_client.my_client.id
}

output "client_secret" {
  value = konnect_identity_auth_server_client.my_client.client_secret
}
