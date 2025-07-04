---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_gateway_plugin_ai_proxy Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  GatewayPluginAiProxy Resource
---

# konnect_gateway_plugin_ai_proxy (Resource)

GatewayPluginAiProxy Resource

## Example Usage

```terraform
resource "konnect_gateway_plugin_ai_proxy" "my_gatewaypluginaiproxy" {
  config = {
    auth = {
      allow_override             = true
      aws_access_key_id          = "...my_aws_access_key_id..."
      aws_secret_access_key      = "...my_aws_secret_access_key..."
      azure_client_id            = "...my_azure_client_id..."
      azure_client_secret        = "...my_azure_client_secret..."
      azure_tenant_id            = "...my_azure_tenant_id..."
      azure_use_managed_identity = true
      gcp_service_account_json   = "...my_gcp_service_account_json..."
      gcp_use_service_account    = false
      header_name                = "...my_header_name..."
      header_value               = "...my_header_value..."
      param_location             = "query"
      param_name                 = "...my_param_name..."
      param_value                = "...my_param_value..."
    }
    llm_format = "bedrock"
    logging = {
      log_payloads   = false
      log_statistics = true
    }
    max_request_body_size = 10
    model = {
      name = "...my_name..."
      options = {
        anthropic_version   = "...my_anthropic_version..."
        azure_api_version   = "...my_azure_api_version..."
        azure_deployment_id = "...my_azure_deployment_id..."
        azure_instance      = "...my_azure_instance..."
        bedrock = {
          aws_assume_role_arn   = "...my_aws_assume_role_arn..."
          aws_region            = "...my_aws_region..."
          aws_role_session_name = "...my_aws_role_session_name..."
          aws_sts_endpoint_url  = "...my_aws_sts_endpoint_url..."
        }
        gemini = {
          api_endpoint = "...my_api_endpoint..."
          location_id  = "...my_location_id..."
          project_id   = "...my_project_id..."
        }
        huggingface = {
          use_cache      = true
          wait_for_model = false
        }
        input_cost     = 7.42
        llama2_format  = "openai"
        max_tokens     = 9
        mistral_format = "ollama"
        output_cost    = 1.81
        temperature    = 2.26
        top_k          = 359
        top_p          = 0.14
        upstream_path  = "...my_upstream_path..."
        upstream_url   = "...my_upstream_url..."
      }
      provider = "anthropic"
    }
    model_name_header  = true
    response_streaming = "allow"
    route_type         = "llm/v1/chat"
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 9
  enabled          = true
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
  updated_at = 3
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `control_plane_id` (String) The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.

### Optional

- `config` (Attributes) (see [below for nested schema](#nestedatt--config))
- `consumer` (Attributes) If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer. (see [below for nested schema](#nestedatt--consumer))
- `consumer_group` (Attributes) If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups (see [below for nested schema](#nestedatt--consumer_group))
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

- `auth` (Attributes) (see [below for nested schema](#nestedatt--config--auth))
- `llm_format` (String) LLM input and output format and schema to use. must be one of ["bedrock", "gemini", "openai"]
- `logging` (Attributes) (see [below for nested schema](#nestedatt--config--logging))
- `max_request_body_size` (Number) max allowed body size allowed to be introspected
- `model` (Attributes) (see [below for nested schema](#nestedatt--config--model))
- `model_name_header` (Boolean) Display the model name selected in the X-Kong-LLM-Model response header
- `response_streaming` (String) Whether to 'optionally allow', 'deny', or 'always' (force) the streaming of answers via server sent events. must be one of ["allow", "always", "deny"]
- `route_type` (String) The model's operation implementation, for this provider. Set to `preserve` to pass through without transformation. must be one of ["llm/v1/chat", "llm/v1/completions", "preserve"]

<a id="nestedatt--config--auth"></a>
### Nested Schema for `config.auth`

Optional:

- `allow_override` (Boolean) If enabled, the authorization header or parameter can be overridden in the request by the value configured in the plugin.
- `aws_access_key_id` (String) Set this if you are using an AWS provider (Bedrock) and you are authenticating using static IAM User credentials. Setting this will override the AWS_ACCESS_KEY_ID environment variable for this plugin instance.
- `aws_secret_access_key` (String) Set this if you are using an AWS provider (Bedrock) and you are authenticating using static IAM User credentials. Setting this will override the AWS_SECRET_ACCESS_KEY environment variable for this plugin instance.
- `azure_client_id` (String) If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the client ID.
- `azure_client_secret` (String) If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the client secret.
- `azure_tenant_id` (String) If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the tenant ID.
- `azure_use_managed_identity` (Boolean) Set true to use the Azure Cloud Managed Identity (or user-assigned identity) to authenticate with Azure-provider models.
- `gcp_service_account_json` (String) Set this field to the full JSON of the GCP service account to authenticate, if required. If null (and gcp_use_service_account is true), Kong will attempt to read from environment variable `GCP_SERVICE_ACCOUNT`.
- `gcp_use_service_account` (Boolean) Use service account auth for GCP-based providers and models.
- `header_name` (String) If AI model requires authentication via Authorization or API key header, specify its name here.
- `header_value` (String) Specify the full auth header value for 'header_name', for example 'Bearer key' or just 'key'.
- `param_location` (String) Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body. must be one of ["body", "query"]
- `param_name` (String) If AI model requires authentication via query parameter, specify its name here.
- `param_value` (String) Specify the full parameter value for 'param_name'.


<a id="nestedatt--config--logging"></a>
### Nested Schema for `config.logging`

Optional:

- `log_payloads` (Boolean) If enabled, will log the request and response body into the Kong log plugin(s) output.
- `log_statistics` (Boolean) If enabled and supported by the driver, will add model usage and token metrics into the Kong log plugin(s) output.


<a id="nestedatt--config--model"></a>
### Nested Schema for `config.model`

Optional:

- `name` (String) Model name to execute.
- `options` (Attributes) Key/value settings for the model (see [below for nested schema](#nestedatt--config--model--options))
- `provider` (String) AI provider request format - Kong translates requests to and from the specified backend compatible formats. must be one of ["anthropic", "azure", "bedrock", "cohere", "gemini", "huggingface", "llama2", "mistral", "openai"]

<a id="nestedatt--config--model--options"></a>
### Nested Schema for `config.model.options`

Optional:

- `anthropic_version` (String) Defines the schema/API version, if using Anthropic provider.
- `azure_api_version` (String) 'api-version' for Azure OpenAI instances.
- `azure_deployment_id` (String) Deployment ID for Azure OpenAI instances.
- `azure_instance` (String) Instance name for Azure OpenAI hosted models.
- `bedrock` (Attributes) (see [below for nested schema](#nestedatt--config--model--options--bedrock))
- `gemini` (Attributes) (see [below for nested schema](#nestedatt--config--model--options--gemini))
- `huggingface` (Attributes) (see [below for nested schema](#nestedatt--config--model--options--huggingface))
- `input_cost` (Number) Defines the cost per 1M tokens in your prompt.
- `llama2_format` (String) If using llama2 provider, select the upstream message format. must be one of ["ollama", "openai", "raw"]
- `max_tokens` (Number) Defines the max_tokens, if using chat or completion models.
- `mistral_format` (String) If using mistral provider, select the upstream message format. must be one of ["ollama", "openai"]
- `output_cost` (Number) Defines the cost per 1M tokens in the output of the AI.
- `temperature` (Number) Defines the matching temperature, if using chat or completion models.
- `top_k` (Number) Defines the top-k most likely tokens, if supported.
- `top_p` (Number) Defines the top-p probability mass, if supported.
- `upstream_path` (String) Manually specify or override the AI operation path, used when e.g. using the 'preserve' route_type.
- `upstream_url` (String) Manually specify or override the full URL to the AI operation endpoints, when calling (self-)hosted models, or for running via a private endpoint.

<a id="nestedatt--config--model--options--bedrock"></a>
### Nested Schema for `config.model.options.bedrock`

Optional:

- `aws_assume_role_arn` (String) If using AWS providers (Bedrock) you can assume a different role after authentication with the current IAM context is successful.
- `aws_region` (String) If using AWS providers (Bedrock) you can override the `AWS_REGION` environment variable by setting this option.
- `aws_role_session_name` (String) If using AWS providers (Bedrock), set the identifier of the assumed role session.
- `aws_sts_endpoint_url` (String) If using AWS providers (Bedrock), override the STS endpoint URL when assuming a different role.


<a id="nestedatt--config--model--options--gemini"></a>
### Nested Schema for `config.model.options.gemini`

Optional:

- `api_endpoint` (String) If running Gemini on Vertex, specify the regional API endpoint (hostname only).
- `location_id` (String) If running Gemini on Vertex, specify the location ID.
- `project_id` (String) If running Gemini on Vertex, specify the project ID.


<a id="nestedatt--config--model--options--huggingface"></a>
### Nested Schema for `config.model.options.huggingface`

Optional:

- `use_cache` (Boolean) Use the cache layer on the inference API
- `wait_for_model` (Boolean) Wait for the model if it is not ready





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
terraform import konnect_gateway_plugin_ai_proxy.my_konnect_gateway_plugin_ai_proxy "{ \"control_plane_id\": \"9524ec7d-36d9-465d-a8c5-83a3c9390458\",  \"id\": \"3473c251-5b6c-4f45-b1ff-7ede735a366d\"}"
```
