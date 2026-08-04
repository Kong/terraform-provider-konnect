resource "konnect_identity_auth_server_client_secret" "my_client_secret" {
  auth_server_id = konnect_identity_auth_server.my_authserver.id
  client_id      = konnect_identity_auth_server_client.my_client.id
  enabled        = true
  secret         = "demo-client-secret"
}
