# Portal Resource
resource "konnect_portal" "my_portal" {
  authentication_enabled    = true
  auto_approve_applications = false
  auto_approve_developers   = false
  default_api_visibility    = "public"
  default_page_visibility   = "private"
  force_destroy             = "true"
  labels = {
    environment = "acceptance-test"
  }
  name         = "TFAcceptancePortalIdpMapping"
  rbac_enabled = true
  sipr_enabled = false
}

# Portal Team
resource "konnect_portal_team" "my_portal_team" {
  name                 = "TFAcceptancePortalTeamGroupMapping"
  description          = "TF acceptance test portal team for group mapping."
  can_own_applications = true
  portal_id            = konnect_portal.my_portal.id
}

# SAML Portal Identity Provider
resource "konnect_portal_identity_provider" "saml_provider" {
  enabled    = true
  type       = "saml"
  login_path = "testportalsamlmapping"
  portal_id  = konnect_portal.my_portal.id
  config = {
    saml_identity_provider_config = {
      idp_metadata_url = "https://mocksaml.com/api/saml/metadata"
    }
  }
}

# Portal IDP Team Group Mapping - SAML
resource "konnect_portal_idp_team_group_mapping" "saml_mapping" {
  group     = "Engineering"
  id        = konnect_portal_identity_provider.saml_provider.id
  portal_id = konnect_portal.my_portal.id
  team_id   = konnect_portal_team.my_portal_team.id
}


# OIDC Portal Identity Provider
resource "konnect_portal_identity_provider" "oidc_provider" {
  enabled    = true
  type       = "oidc"
  login_path = "testportaloidcmapping"
  portal_id  = konnect_portal.my_portal.id
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

# Portal IDP Team Group Mapping - OIDC
resource "konnect_portal_idp_team_group_mapping" "oidc_mapping" {
  group     = "Tech Leads"
  id        = konnect_portal_identity_provider.oidc_provider.id
  portal_id = konnect_portal.my_portal.id
  team_id   = konnect_portal_team.my_portal_team.id
}

