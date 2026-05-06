resource "konnect_team" "my_team" {
  name        = "TFAcceptanceTeamGroupMapping"
  description = "TF acceptance test team for group mapping."
}

# SAML Identity Provider
resource "konnect_identity_provider" "saml_provider" {
  enabled    = true
  type       = "saml"
  login_path = "testsamlmapping"
  config = {
    saml_identity_provider_config = {
      idp_metadata_url = "https://mocksaml.com/api/saml/metadata"
    }
  }
}

# Identity Provider Team Group Mapping - SAML
resource "konnect_identity_provider_team_group_mapping" "my_mapping" {
  group                = "Engineering"
  identity_provider_id = konnect_identity_provider.saml_provider.id
  team_id              = konnect_team.my_team.id
}


# OIDC Identity Provider
resource "konnect_identity_provider" "oidc_provider" {
  enabled    = true
  type       = "oidc"
  login_path = "testoidcmapping"
  config = {
    oidc_identity_provider_config = {
      client_id     = "client-id"
      client_secret = "my-client-secret"
      issuer_url    = "https://accounts.google.com"
      scopes = [
        "openid",
        "email",
        "profile"
      ]
      claim_mappings = {
        email  = "email"
        groups = "groups"
        name   = "name"
      }
    }
  }
}

# Identity Provider Team Group Mapping - OIDC
resource "konnect_identity_provider_team_group_mapping" "my_mapping" {
  group                = "Tech Leads"
  identity_provider_id = konnect_identity_provider.oidc_provider.id
  team_id              = konnect_team.my_team.id
}