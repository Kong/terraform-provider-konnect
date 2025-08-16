resource "konnect_portal" "my_portal_for_customization" {
  force_destroy = "true"
  name         = "My v3 portal name"
}

resource "konnect_portal_customization" "my_portal_customization" {
  portal_id = konnect_portal.my_portal_for_customization.id
  robots    = "my_robots"
  theme = {
    colors = {
      primary = "#000000"
    }
    mode = "system"
    name = "my_name"
  }
}