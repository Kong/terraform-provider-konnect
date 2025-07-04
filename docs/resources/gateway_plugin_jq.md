---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_gateway_plugin_jq Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  GatewayPluginJq Resource
---

# konnect_gateway_plugin_jq (Resource)

GatewayPluginJq Resource

## Example Usage

```terraform
resource "konnect_gateway_plugin_jq" "my_gatewaypluginjq" {
  config = {
    request_if_media_type = [
      "..."
    ]
    request_jq_program = "...my_request_jq_program..."
    request_jq_program_options = {
      ascii_output   = false
      compact_output = false
      join_output    = true
      raw_output     = true
      sort_keys      = false
    }
    response_if_media_type = [
      "..."
    ]
    response_if_status_code = [
      188
    ]
    response_jq_program = "...my_response_jq_program..."
    response_jq_program_options = {
      ascii_output   = false
      compact_output = true
      join_output    = true
      raw_output     = false
      sort_keys      = false
    }
  }
  consumer = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 5
  enabled          = false
  id               = "...my_id..."
  instance_name    = "...my_instance_name..."
  ordering = {
    after = {
      access = [
        "..."
      ]
    }
    before = {
      access = [
        "..."
      ]
    }
  }
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
  protocols = [
    "grpcs"
  ]
  route = {
    id = "...my_id..."
  }
  service = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
  updated_at = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `control_plane_id` (String) The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.

### Optional

- `config` (Attributes) (see [below for nested schema](#nestedatt--config))
- `consumer` (Attributes) If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer. (see [below for nested schema](#nestedatt--consumer))
- `created_at` (Number) Unix epoch when the resource was created.
- `enabled` (Boolean) Whether the plugin is applied.
- `instance_name` (String)
- `ordering` (Attributes) (see [below for nested schema](#nestedatt--ordering))
- `partials` (Attributes List) (see [below for nested schema](#nestedatt--partials))
- `protocols` (Set of String) A set of strings representing HTTP protocols.
- `route` (Attributes) If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used. (see [below for nested schema](#nestedatt--route))
- `service` (Attributes) If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched. (see [below for nested schema](#nestedatt--service))
- `tags` (List of String) An optional set of strings associated with the Plugin for grouping and filtering.
- `updated_at` (Number) Unix epoch when the resource was last updated.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `request_if_media_type` (List of String)
- `request_jq_program` (String)
- `request_jq_program_options` (Attributes) (see [below for nested schema](#nestedatt--config--request_jq_program_options))
- `response_if_media_type` (List of String)
- `response_if_status_code` (List of Number)
- `response_jq_program` (String)
- `response_jq_program_options` (Attributes) (see [below for nested schema](#nestedatt--config--response_jq_program_options))

<a id="nestedatt--config--request_jq_program_options"></a>
### Nested Schema for `config.request_jq_program_options`

Optional:

- `ascii_output` (Boolean)
- `compact_output` (Boolean)
- `join_output` (Boolean)
- `raw_output` (Boolean)
- `sort_keys` (Boolean)


<a id="nestedatt--config--response_jq_program_options"></a>
### Nested Schema for `config.response_jq_program_options`

Optional:

- `ascii_output` (Boolean)
- `compact_output` (Boolean)
- `join_output` (Boolean)
- `raw_output` (Boolean)
- `sort_keys` (Boolean)



<a id="nestedatt--consumer"></a>
### Nested Schema for `consumer`

Optional:

- `id` (String)


<a id="nestedatt--ordering"></a>
### Nested Schema for `ordering`

Optional:

- `after` (Attributes) (see [below for nested schema](#nestedatt--ordering--after))
- `before` (Attributes) (see [below for nested schema](#nestedatt--ordering--before))

<a id="nestedatt--ordering--after"></a>
### Nested Schema for `ordering.after`

Optional:

- `access` (List of String)


<a id="nestedatt--ordering--before"></a>
### Nested Schema for `ordering.before`

Optional:

- `access` (List of String)



<a id="nestedatt--partials"></a>
### Nested Schema for `partials`

Optional:

- `id` (String)
- `name` (String)
- `path` (String)


<a id="nestedatt--route"></a>
### Nested Schema for `route`

Optional:

- `id` (String)


<a id="nestedatt--service"></a>
### Nested Schema for `service`

Optional:

- `id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_gateway_plugin_jq.my_konnect_gateway_plugin_jq "{ \"control_plane_id\": \"9524ec7d-36d9-465d-a8c5-83a3c9390458\",  \"id\": \"3473c251-5b6c-4f45-b1ff-7ede735a366d\"}"
```
