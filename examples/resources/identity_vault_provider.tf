# A static_secret provider stores a single pre-issued credential, such as an API
# key. The credential itself is a konnect_identity_vault_provider_credential;
# client_id and client_secret are rejected for this template.
resource "konnect_identity_vault_provider" "my_static_secret_provider" {
  directory_id    = konnect_identity_directory.my_directory.id
  name            = "my-static-secret-provider"
  template_name   = "static_secret"
  credential_type = "shared"
}

# An OAuth template. credential_type = "user" has every principal enroll its own
# upstream credential, "shared" uses one credential for the whole directory.
resource "konnect_identity_vault_provider" "my_github_provider" {
  directory_id    = konnect_identity_directory.my_directory.id
  name            = "my-github-provider"
  template_name   = "github"
  credential_type = "user"

  client_id     = "demo-github-client-id"
  client_secret = "demo-github-client-secret"

  scopes = ["read:user", "repo"]
}
