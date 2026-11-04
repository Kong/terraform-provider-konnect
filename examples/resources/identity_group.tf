resource "konnect_identity_group" "my_group" {
  directory_id = konnect_identity_directory.my_directory.id
  name         = "paymentsTeam"
  description  = "Team that owns the payments platform"

  # "group" is a plain group, "team" carries a team_manager identity.
  type = "group"

  # Metadata is attached to the group and passed to the data plane as part of
  # the execution context.
  metadata = jsonencode({
    cost_center = "payments"
  })
}
