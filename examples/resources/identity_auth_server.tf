resource "konnect_identity_auth_server" "my_authserver" {
  name     = "my-auth-server"
  audience = "local-demo"
}
