# type = "v1" lets Konnect generate the key. The generated value appears only in
# the create response, not in state, so this resource cannot read it back. Use
# type = "imported" to supply your own secret.
resource "konnect_identity_principal_api_key" "my_principal_api_key" {
  directory_id = konnect_identity_directory.my_directory.id
  principal_id = konnect_identity_principal.my_principal.id
  name         = "payments-service-api-key"
  type         = "v1"
}

# type = "imported" registers a key that already exists outside Konnect.
resource "konnect_identity_principal_api_key" "my_imported_principal_api_key" {
  directory_id = konnect_identity_directory.my_directory.id
  principal_id = konnect_identity_principal.my_principal.id
  name         = "payments-service-imported-key"
  type         = "imported"
  secret       = "demo-imported-principal-api-key"
}
