resource "konnect_system_account" "my_systemaccount" {
  description     = "my description"
  konnect_managed = false
  name            = "TFAcceptanceSystemAccountName"
}

resource "konnect_team" "my_team" {
  name               = "TFAcceptanceSysAccTeamName"
  description        = "my team description"
}

resource "konnect_system_account_team" "my_systemaccount_team" {
  account_id = konnect_system_account.my_systemaccount.id
  team_id           = konnect_team.my_team.id
}