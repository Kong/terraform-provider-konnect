resource "konnect_system_account_team" "my_systemaccountteam" {
  account_id = konnect_system_account.my_systemaccount.id

  # Update this to your team ID manually until the Teams data source is available
  # You can find this using the API
  # GET https://global.api.konghq.com/v3/teams
  team_id = "3e3a7539-4d61-4424-a714-fdfdc3728fad"
}
