resource "konnect_portal" "my_portal_for_favicon" {
  force_destroy = "true"
  name         = "My v3 portal name"
}

resource "konnect_portal_favicon" "my_portal_favicon" {
  data      = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR4nGMAAQAABQABDQottAAAAABJRU5ErkJggg=="
  portal_id = konnect_portal.my_portal_for_favicon.id
}