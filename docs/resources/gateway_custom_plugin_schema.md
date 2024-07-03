---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_gateway_custom_plugin_schema Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  GatewayCustomPluginSchema Resource
---

# konnect_gateway_custom_plugin_schema (Resource)

GatewayCustomPluginSchema Resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `control_plane_id` (String) The UUID of your control plane. This variable is available in the Konnect manager
- `lua_schema` (String) The custom plugin schema; `jq -Rs '.' schema.lua`.

### Read-Only

- `created_at` (Number) An ISO-8604 timestamp representation of custom plugin schema creation date.
- `name` (String) The custom plugin name
- `updated_at` (Number) An ISO-8604 timestamp representation of custom plugin schema update date.

