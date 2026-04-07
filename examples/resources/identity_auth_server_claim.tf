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
