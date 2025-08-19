resource "konnect_portal" "my_portal" {
  force_destroy                        = "true"
  default_api_visibility               = "public"
  name         = "My Developer Portal"
}

resource "konnect_portal_snippet" "my_portal_snippet" {
  content     = "# Welcome to My Snippet"
  description = "A custom snippet about developer portals"
  name        = "my-snippet"
  title       = "My Snippet"
  portal_id   = konnect_portal.my_portal.id
}

resource "konnect_portal_page" "my_portalpage" {
  content        = "# Welcome to My First Page"
  description    = "A custom page about my developer portal"
  portal_id      = konnect_portal.my_portal.id
  slug           = "/my-first-page"
  title          = "My First Page"
  status         = "published"
  visibility     = "public"
}

resource "konnect_portal_logo" "my_portal_logo" {
  data      = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR4nGMAAQAABQABDQottAAAAABJRU5ErkJggg=="
  portal_id = konnect_portal.my_portal.id
}

resource "konnect_portal_favicon" "my_portal_favicon" {
  data      = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR4nGMAAQAABQABDQottAAAAABJRU5ErkJggg=="
  portal_id = konnect_portal.my_portal.id
}

resource "konnect_portal_customization" "my_portal_customization" {
  portal_id = konnect_portal.my_portal.id
  robots    = "my_robots_txt"
  theme = {
    colors = {
      primary = "#000000"
    }
    mode = "system"
    name = "my_system_theme"
  }
}

resource "konnect_portal_custom_domain" "my_portal_custom_domain" {
  enabled   = false
  hostname  = "my.custom.domain"
  portal_id = konnect_portal.my_portal.id
  ssl = {
    domain_verification_method = "http"
  }
}
