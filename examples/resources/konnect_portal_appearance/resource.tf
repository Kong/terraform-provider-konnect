resource "konnect_portal_appearance" "my_portalappearance" {
  custom_fonts = {
    base     = "Roboto"
    code     = "Roboto"
    headings = "Roboto"
  }
  custom_theme = {
    colors = {
      button = {
        primary_fill = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        primary_text = {
          description = "...my_description..."
          value       = "...my_value..."
        }
      }
      section = {
        accent = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        body = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        footer = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        header = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        hero = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        stroke = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        tertiary = {
          description = "...my_description..."
          value       = "...my_value..."
        }
      }
      text = {
        accent = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        footer = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        header = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        headings = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        hero = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        link = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        primary = {
          description = "...my_description..."
          value       = "...my_value..."
        }
        secondary = {
          description = "...my_description..."
          value       = "...my_value..."
        }
      }
    }
  }
  images = {
    catalog_cover = {
      data     = "data:image/png,YW5faW1hZ2VfZmlsZQ=="
      filename = "...my_filename..."
    }
    favicon = {
      data     = "data:image/png,YW5faW1hZ2VfZmlsZQ=="
      filename = "...my_filename..."
    }
    logo = {
      data     = "data:image/png,YW5faW1hZ2VfZmlsZQ=="
      filename = "...my_filename..."
    }
  }
  portal_id = "3d7b380a-055a-48ae-88ff-cc687bbb2679"
  text = {
    catalog = {
      primary_header  = "...my_primary_header..."
      welcome_message = "...my_welcome_message..."
    }
  }
  theme_name       = "mint_rocket"
  use_custom_fonts = false
}