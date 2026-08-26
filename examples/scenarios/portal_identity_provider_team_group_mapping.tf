resource "konnect_portal" "my_portal" {
  name          = "example-portal"
}

resource "konnect_portal_team" "my_team" {
  portal_id = konnect_portal.my_portal.id
  name        = "example-team"
  description = "Sample description"
}

resource "konnect_portal_identity_provider" "saml_provider" {
  type = "saml"
  enabled = true
  portal_id = konnect_portal.my_portal.id
  config = {
    saml_identity_provider_config = {
      idp_metadata_url = "https://mocksaml.com/api/saml/metadata" // //file("./saml-metadata.xml")
      idp_metadata_xml = ""
    }
  }
}

resource "konnect_portal_identity_provider_team_group_mapping" "my_portalidentityproviderteamgroupmapping" {
  group     = "example-group"
  identity_provider_id        = konnect_portal_identity_provider.saml_provider.id
  portal_id = konnect_portal.my_portal.id
  team_id   = konnect_portal_team.my_team.id
}
