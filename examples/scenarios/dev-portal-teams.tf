# Create a new Dev Portal
resource "konnect_portal" "my_portal" {
  force_destroy                        = "true"
  name         = "My v3 portal name"
}
# Create a new team in the Dev Portal
resource "konnect_portal_team" "myteam" {
    portal_id = konnect_portal.my_portal.id
    name = "My platform team"
    description = "Description for my platform team"
}