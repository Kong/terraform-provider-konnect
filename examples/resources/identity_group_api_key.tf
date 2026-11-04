# type = "v1" lets Konnect generate the key. The generated value appears only in
# the create response, not in state, so this resource cannot read it back. Use
# type = "imported" to supply your own secret.
resource "konnect_identity_group_api_key" "my_group_api_key" {
  directory_id = konnect_identity_directory.my_directory.id
  group_id     = konnect_identity_group.my_group.id
  name         = "payments-team-api-key"
  type         = "v1"
}

# type = "imported" registers a key that already exists outside Konnect.
resource "konnect_identity_group_api_key" "my_imported_group_api_key" {
  directory_id = konnect_identity_directory.my_directory.id
  group_id     = konnect_identity_group.my_group.id
  name         = "payments-team-imported-key"
  type         = "imported"
  secret       = "demo-imported-group-api-key"
}
