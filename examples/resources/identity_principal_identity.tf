# Set one of custom, oidc, auth_server_client or control_plane_consumer per
# resource to link the principal to an identity it authenticates with.

# A custom key/value identity, for example an email address.
resource "konnect_identity_principal_identity" "my_principal_email" {
  directory_id = konnect_identity_directory.my_directory.id
  principal_id = konnect_identity_principal.my_principal.id

  custom = {
    key   = "email"
    value = "payments@example.com"
  }
}

# An issuer and claim pair, for example when a client presents an OIDC token.
resource "konnect_identity_principal_identity" "my_principal_oidc" {
  directory_id = konnect_identity_directory.my_directory.id
  principal_id = konnect_identity_principal.my_principal.id

  oidc = {
    issuer = "https://accounts.example.com"

    claim = {
      name  = "sub"
      value = "payments-service"
    }
  }
}
