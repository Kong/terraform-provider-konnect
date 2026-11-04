# The username half of a basic auth credential. Create the password with
# konnect_identity_principal_basic_auth_password.
resource "konnect_identity_principal_basic_auth" "my_basic_auth" {
  directory_id = konnect_identity_directory.my_directory.id
  principal_id = konnect_identity_principal.my_principal.id
  username     = "payments-service"
}
