# Adds a principal to a group. Set member.type to "group" to nest a group
# inside this group instead.
resource "konnect_identity_group_membership" "my_group_membership" {
  directory_id = konnect_identity_directory.my_directory.id
  group_id     = konnect_identity_group.my_group.id

  member = {
    id   = konnect_identity_principal.my_principal.id
    type = "principal"
  }
}
