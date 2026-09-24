# Create Team
resource "konnect_team" "my_team" {
  name        = "TFAcceptanceTeamGroupMappingSAML"
  description = "TF acceptance test team for SAML group mapping."
}

# SAML Identity Provider
resource "konnect_identity_provider" "saml_provider" {
  enabled    = false // to avoid race condition - konnect only allows one enabled idp at a time.
  type       = "saml"
  login_path = "testsamlmapping"
  config = {
    saml_identity_provider_config = {
      idp_metadata_url = "https://mocksaml.com/api/saml/metadata"
    }
  }
}

# Identity Provider Team Group Mapping
resource "konnect_identity_provider_team_group_mapping" "my_mapping" {
  group                = "Engineering"
  identity_provider_id = konnect_identity_provider.saml_provider.id
  team_id              = konnect_team.my_team.id
}

# Identity Provider datasource.
data "konnect_identity_provider" "my_identity_provider" {
  id = konnect_identity_provider.saml_provider.id
}


// Identity Provider List datasource.
data "konnect_identity_provider_list" "my_list_identity_provider" {
  filter = {
    type = {
      eq = "saml"
    }
  }

  depends_on = [konnect_identity_provider.saml_provider]
}

output "identity_provider_list" {
  value = data.konnect_identity_provider_list.my_list_identity_provider
}

output "identity_provider" {
  value = data.konnect_identity_provider.my_identity_provider
}