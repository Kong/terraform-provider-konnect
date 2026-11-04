# The default credential of a static_secret provider: no selector, so it applies
# to every caller that the provider serves.
resource "konnect_identity_vault_provider_credential" "my_default_credential" {
  directory_id = konnect_identity_directory.my_directory.id
  provider_id  = konnect_identity_vault_provider.my_static_secret_provider.id
  secret       = "demo-static-secret"
}

# A selector narrows the credential to the callers its CEL expression matches.
# priority orders the selectors, lowest first, and is required with a selector.
resource "konnect_identity_vault_provider_credential" "my_payments_credential" {
  directory_id = konnect_identity_directory.my_directory.id
  provider_id  = konnect_identity_vault_provider.my_static_secret_provider.id
  priority     = 1
  selector     = "consumer.name == \"payments-service\""
  secret       = "demo-payments-secret"
}
