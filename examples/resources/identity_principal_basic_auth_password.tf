# The password half of a basic auth credential. The secret is returned only on
# create, so keep it wherever the client reads its credentials from.
resource "konnect_identity_principal_basic_auth_password" "my_basic_auth_password" {
  directory_id  = konnect_identity_directory.my_directory.id
  principal_id  = konnect_identity_principal.my_principal.id
  basic_auth_id = konnect_identity_principal_basic_auth.my_basic_auth.id
  secret        = "demo-basic-auth-password"
}
