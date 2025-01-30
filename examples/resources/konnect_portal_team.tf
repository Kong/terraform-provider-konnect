resource "konnect_portal_team" "myteam" {
    portal_id = konnect_portal.my_portal.id
    name = "My platform team"
    description = "Description for my platform team"
}