---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_gateway_plugin_http_log Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  GatewayPluginHTTPLog Resource
---

# konnect_gateway_plugin_http_log (Resource)

GatewayPluginHTTPLog Resource

## Example Usage

```terraform
resource "konnect_gateway_plugin_http_log" "my_gatewaypluginhttplog" {
  config = {
    content_type = "application/json; charset=utf-8"
    custom_fields_by_lua = {
      "see" : jsonencode("documentation"),
    }
    flush_timeout = 8.97
    headers = {
      "see" : jsonencode("documentation"),
    }
    http_endpoint = "...my_http_endpoint..."
    keepalive     = 2.25
    method        = "PUT"
    queue = {
      concurrency_limit    = 0
      initial_retry_delay  = 154435.3
      max_batch_size       = 297113
      max_bytes            = 7
      max_coalescing_delay = 198.88
      max_entries          = 603921
      max_retry_delay      = 290740.88
      max_retry_time       = 4.99
    }
    queue_size  = 3
    retry_count = 10
    timeout     = 6.49
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = false
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
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `control_plane_id` (String) The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.

### Optional

- `config` (Attributes) (see [below for nested schema](#nestedatt--config))
- `consumer` (Attributes) If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer. (see [below for nested schema](#nestedatt--consumer))
- `consumer_group` (Attributes) (see [below for nested schema](#nestedatt--consumer_group))
- `enabled` (Boolean) Whether the plugin is applied.
- `instance_name` (String)
- `ordering` (Attributes) (see [below for nested schema](#nestedatt--ordering))
- `protocols` (List of String) A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
- `route` (Attributes) If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used. (see [below for nested schema](#nestedatt--route))
- `service` (Attributes) If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched. (see [below for nested schema](#nestedatt--service))
- `tags` (List of String) An optional set of strings associated with the Plugin for grouping and filtering.

### Read-Only

- `created_at` (Number) Unix epoch when the resource was created.
- `id` (String) The ID of this resource.
- `updated_at` (Number) Unix epoch when the resource was last updated.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `content_type` (String) Indicates the type of data sent. The only available option is `application/json`. must be one of ["application/json", "application/json; charset=utf-8"]
- `custom_fields_by_lua` (Map of String) Lua code as a key-value map
- `flush_timeout` (Number) Optional time in seconds. If `queue_size` > 1, this is the max idle time before sending a log with less than `queue_size` records.
- `headers` (Map of String) An optional table of headers included in the HTTP message to the upstream server. Values are indexed by header name, and each header name accepts a single string.
- `http_endpoint` (String) A string representing a URL, such as https://example.com/path/to/resource?q=search.
- `keepalive` (Number) An optional value in milliseconds that defines how long an idle connection will live before being closed.
- `method` (String) An optional method used to send data to the HTTP server. Supported values are `POST` (default), `PUT`, and `PATCH`. must be one of ["POST", "PUT", "PATCH"]
- `queue` (Attributes) (see [below for nested schema](#nestedatt--config--queue))
- `queue_size` (Number) Maximum number of log entries to be sent on each message to the upstream server.
- `retry_count` (Number) Number of times to retry when sending data to the upstream server.
- `timeout` (Number) An optional timeout in milliseconds when sending data to the upstream server.

<a id="nestedatt--config--queue"></a>
### Nested Schema for `config.queue`

Optional:

- `concurrency_limit` (Number) The number of of queue delivery timers. -1 indicates unlimited. must be one of ["-1", "1"]
- `initial_retry_delay` (Number) Time in seconds before the initial retry is made for a failing batch.
- `max_batch_size` (Number) Maximum number of entries that can be processed at a time.
- `max_bytes` (Number) Maximum number of bytes that can be waiting on a queue, requires string content.
- `max_coalescing_delay` (Number) Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.
- `max_entries` (Number) Maximum number of entries that can be waiting on the queue.
- `max_retry_delay` (Number) Maximum time in seconds between retries, caps exponential backoff.
- `max_retry_time` (Number) Time in seconds before the queue gives up calling a failed handler for a batch.



<a id="nestedatt--consumer"></a>
### Nested Schema for `consumer`

Optional:

- `id` (String)


<a id="nestedatt--consumer_group"></a>
### Nested Schema for `consumer_group`

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
terraform import konnect_gateway_plugin_http_log.my_konnect_gateway_plugin_http_log "{ \"control_plane_id\": \"9524ec7d-36d9-465d-a8c5-83a3c9390458\",  \"plugin_id\": \"3473c251-5b6c-4f45-b1ff-7ede735a366d\"}"
```