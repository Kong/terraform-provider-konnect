resource "konnect_portal_customization" "my_portalcustomization" {
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
        title = "My Title"
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
  portal_id = konnect_portal.my_portal.id
  spec_renderer = {
    hide_deprecated = true
    hide_internal   = false
    infinite_scroll = false
    show_schemas    = false
    try_it_insomnia = false
    try_it_ui       = false
  }
  theme = {
    colors = {
      primary = "#000000"
    }
    mode = "system"
    name = "My Theme"
  }
}