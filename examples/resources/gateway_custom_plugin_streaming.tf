resource "konnect_gateway_control_plane" "tfdemo" {
  name = "Terraform Control Plane"
}

# Define a custom Lua plugin with handler and schema
resource "konnect_gateway_custom_plugin_streaming" "header_injector" {
  name = "setheader-streaming"

  handler = <<-EOT
    return {
      VERSION = "1.0.0",
      PRIORITY = 500,
      access = function(self, config)
        kong.service.request.set_header(config.name, config.value)
      end
    }
  EOT

  schema = <<-EOT
    return {
      name = "setheader-streaming",
      fields = {
        { protocols = require("kong.db.schema.typedefs").protocols_http },
        {
          config = {
            type = "record",
            fields = {
              { name = { description = "The name of the header to set.", type = "string", required = true, }, },
              { value = { description = "The value for the header.", type = "string", required = true, }, },
            },
          },
        },
      }
    }
  EOT

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Create an instance of the custom plugin
resource "konnect_gateway_custom_plugin" "inject_header" {
  name          = "setheader-streaming"
  instance_name = "inject-custom-header"

  config = {
    name  = "x-custom-header"
    value = "my-custom-value"
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id

  depends_on = [konnect_gateway_custom_plugin_streaming.header_injector]
}
