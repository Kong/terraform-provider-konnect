resource "konnect_gateway_custom_plugin_schema" "my_gatewaycustompluginschema" {
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  lua_schema       = "return { name = \\\"myplugin\\\", fields = { { config = { type = \\\"record\\\", fields = { } } } } }"
}