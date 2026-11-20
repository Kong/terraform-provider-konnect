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
  css = "my_css"
  menu = {
    footer_bottom = [
      {
        external   = false
        path       = "/about/company"
        title      = "My Page"
        visibility = "public"
      }
    ]
    footer_sections = [
      {
        items = [
          {
            external   = true
            path       = "/about/company"
            title      = "My Page"
            visibility = "public"
          }
        ]
        title = "...my_title..."
      }
    ]
    main = [
      {
        external   = true
        path       = "/about/company"
        title      = "My Page"
        visibility = "public"
      }
    ]
  }
  spec_renderer = {
    infinite_scroll = true
    show_schemas    = false
    try_it_insomnia = false
    try_it_ui       = true
  }
}