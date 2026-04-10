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
