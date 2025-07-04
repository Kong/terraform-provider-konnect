---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_gateway_plugin_request_validator Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  GatewayPluginRequestValidator Resource
---

# konnect_gateway_plugin_request_validator (Resource)

GatewayPluginRequestValidator Resource

## Example Usage

```terraform
resource "konnect_gateway_plugin_request_validator" "my_gatewaypluginrequestvalidator" {
  config = {
    allowed_content_types = [
      "..."
    ]
    body_schema                       = "...my_body_schema..."
    content_type_parameter_validation = false
    parameter_schema = [
      {
        explode  = true
        in       = "query"
        name     = "...my_name..."
        required = false
        schema   = "...my_schema..."
        style    = "matrix"
      }
    ]
    verbose_response = true
    version          = "kong"
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
  updated_at = 7
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

- `allowed_content_types` (List of String) List of allowed content types. The value can be configured with the `charset` parameter. For example, `application/json; charset=UTF-8`.
- `body_schema` (String) The request body schema specification. One of `body_schema` or `parameter_schema` must be specified.
- `content_type_parameter_validation` (Boolean) Determines whether to enable parameters validation of request content-type.
- `parameter_schema` (Attributes List) Array of parameter validator specification. One of `body_schema` or `parameter_schema` must be specified. (see [below for nested schema](#nestedatt--config--parameter_schema))
- `verbose_response` (Boolean) If enabled, the plugin returns more verbose and detailed validation errors.
- `version` (String) Which validator to use. Supported values are `kong` (default) for using Kong's own schema validator, or `draft4` for using a JSON Schema Draft 4-compliant validator. must be one of ["draft4", "kong"]

<a id="nestedatt--config--parameter_schema"></a>
### Nested Schema for `config.parameter_schema`

Optional:

- `explode` (Boolean) Required when `schema` and `style` are set. When `explode` is `true`, parameter values of type `array` or `object` generate separate parameters for each value of the array or key-value pair of the map. For other types of parameters, this property has no effect.
- `in` (String) The location of the parameter. Not Null; must be one of ["header", "path", "query"]
- `name` (String) The name of the parameter. Parameter names are case-sensitive, and correspond to the parameter name used by the `in` property. If `in` is `path`, the `name` field MUST correspond to the named capture group from the configured `route`. Not Null
- `required` (Boolean) Determines whether this parameter is mandatory. Not Null
- `schema` (String) Requred when `style` and `explode` are set. This is the schema defining the type used for the parameter. It is validated using `draft4` for JSON Schema draft 4 compliant validator. In addition to being a valid JSON Schema, the parameter schema MUST have a top-level `type` property to enable proper deserialization before validating.
- `style` (String) Required when `schema` and `explode` are set. Describes how the parameter value will be deserialized depending on the type of the parameter value. must be one of ["deepObject", "form", "label", "matrix", "pipeDelimited", "simple", "spaceDelimited"]



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
terraform import konnect_gateway_plugin_request_validator.my_konnect_gateway_plugin_request_validator "{ \"control_plane_id\": \"9524ec7d-36d9-465d-a8c5-83a3c9390458\",  \"id\": \"3473c251-5b6c-4f45-b1ff-7ede735a366d\"}"
```
