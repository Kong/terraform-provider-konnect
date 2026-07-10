resource "konnect_organization_personal_access_token_settings" "my_organizationpersonalaccesstokensettings" {
  max_expiration_period_days = 365
  organization_id            = "72ab1dcd-3e4a-4f9b-9e9d-94217898055a"
  pats_enabled               = true
}
