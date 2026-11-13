resource "konnect_portal" "my_portal_for_logo" {
  force_destroy = "true"
  name         = "My v3 portal name"
}

resource "konnect_portal_logo" "my_portal_logo" {
  data      = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGL6n50NCAAA//8ESgHYqxSkJQAAAABJRU5ErkJggg=="
  portal_id = konnect_portal.my_portal_for_logo.id
}