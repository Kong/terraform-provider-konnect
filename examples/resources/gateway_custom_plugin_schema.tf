resource "konnect_gateway_custom_plugin_schema" "foo" {
  lua_schema = <<EOF
return {
  name = "myplugin", 
  fields = { 
    {
      config = { 
        type = "record",
        fields = {
          { hello = { type = "string", required = true } },
        } 
      }
    } 
  }
}
EOF

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
