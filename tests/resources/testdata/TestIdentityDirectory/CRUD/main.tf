resource "konnect_identity_directory" "my_identitydirectory" {
  name          = "tf-test-identity-directory"
  description   = "TF test directory"
  force_destroy = "true"
  vault_enabled = true
}

resource "konnect_identity_group" "my_identitygroup" {
  directory_id = konnect_identity_directory.my_identitydirectory.id
  name         = "tftestgroup"
  description  = "TF test group"
}

resource "konnect_identity_principal" "my_identityprincipal" {
  directory_id = konnect_identity_directory.my_identitydirectory.id
  display_name = "TF Test Principal"
  description  = "TF test principal"
}

resource "konnect_identity_group_membership" "my_identitygroupmembership" {
  directory_id = konnect_identity_directory.my_identitydirectory.id
  group_id     = konnect_identity_group.my_identitygroup.id
  member = {
    id   = konnect_identity_principal.my_identityprincipal.id
    type = "principal"
  }
}

resource "konnect_identity_group_api_key" "my_identitygroupapikey" {
  directory_id = konnect_identity_directory.my_identitydirectory.id
  group_id     = konnect_identity_group.my_identitygroup.id
  name = "tftest-group-key"
  type = "v1"
}

resource "konnect_identity_principal_api_key" "my_identityprincipalapikey" {
  directory_id = konnect_identity_directory.my_identitydirectory.id
  principal_id = konnect_identity_principal.my_identityprincipal.id
  name = "tftest-principal-key"
  type = "v1"
}

resource "konnect_identity_principal_basic_auth" "my_identityprincipalbasicauth" {
  directory_id = konnect_identity_directory.my_identitydirectory.id
  principal_id = konnect_identity_principal.my_identityprincipal.id
  username     = "tftest-user"
}

resource "konnect_identity_principal_basic_auth_password" "my_identityprincipalbasicauthpassword" {
  basic_auth_id = konnect_identity_principal_basic_auth.my_identityprincipalbasicauth.id
  directory_id  = konnect_identity_directory.my_identitydirectory.id
  principal_id  = konnect_identity_principal.my_identityprincipal.id
  secret        = "tftest-basic-auth-secret"
}

resource "konnect_identity_principal_identity" "my_identityprincipalidentity" {
  directory_id = konnect_identity_directory.my_identitydirectory.id
  principal_id = konnect_identity_principal.my_identityprincipal.id
  custom = {
    key   = "tftest-key"
    value = "tftest-value"
  }
}

resource "konnect_identity_trusted_idp" "my_identitytrustedidp" {
  directory_id = konnect_identity_directory.my_identitydirectory.id
  display_name = "TF Test Trusted IdP"
  issuer_url   = "https://accounts.google.com"
  //jwks_uri     = "https://accounts.google.com/.well-known/jwks.json"
}

resource "konnect_identity_vault_key" "my_identityvaultkey" {
  directory_id = konnect_identity_directory.my_identitydirectory.id
  name         = "tftest-vault-key"
  public_key   = <<-EOT
    -----BEGIN PUBLIC KEY-----
    MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6eLB+91ZbuIxBLwkVH79
    MGgWDT43PA10ONxcu3k6lBerKJQ9wkO3kHqfa6ASoh1HaVTlxgRckD+KldCIfOOF
    F7upQhpqlasl0rqxT/ahGPtdudVyHpVfprR7qg5izk+VBqqQOKoIEmX/EPGUuy/Z
    ASxnQEWSCFNTL+h6TPc1lKK8kei+CsuIj8r+Zgrh32eJ/cxHMzAE0lokt0F84W/8
    VSNiOqoKblsfBBFaK91CVdS5IHfHbIdBKWVjEvCo+Ve4NKrGoLtEbfemIXn1PrlM
    dAQq/Uw+n39aOprLP/qZAgtzDrXy/6TtigcGrSlSRiFSvZBizbmWfJIysJcEBGmj
    twIDAQAB
    -----END PUBLIC KEY-----
  EOT
}

resource "konnect_identity_vault_provider" "my_identityvaultprovider" {
  directory_id    = konnect_identity_directory.my_identitydirectory.id
  name            = "tftest-vault-provider"
  template_name   = "static_secret"
  credential_type = "shared"
}

resource "konnect_identity_vault_provider_credential" "my_identityvaultprovidercredential" {
  directory_id = konnect_identity_directory.my_identitydirectory.id
  provider_id  = konnect_identity_vault_provider.my_identityvaultprovider.id
  secret       = "tftest-static-secret"
}
