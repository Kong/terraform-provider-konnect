# Create Team
resource "konnect_team" "my_team" {
  name        = "TFAcceptanceTeamGroupMappingSAML"
  description = "TF acceptance test team for SAML group mapping."
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
  lifecycle {
    ignore_changes = [config.saml_identity_provider_config.idp_metadata_xml]
  }
}

# Identity Provider Team Group Mapping
resource "konnect_identity_provider_team_group_mapping" "my_mapping" {
  group                = "Engineering"
  identity_provider_id = konnect_identity_provider.saml_provider.id
  team_id              = konnect_team.my_team.id
}