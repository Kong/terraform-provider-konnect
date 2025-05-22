# Changelog

## 2.8.1  
> Released on 2025/05/22

### Bug fixes

* The provider now works with terraform v1.12.1

## 2.8.0  
> Released on 2025/05/20

### Features

* Add support for new resource `konnect_gateway_partial` which enables users to define shared configuration for Redis.

## 2.7.4
> Released on 2025/05/14

### Bug fixes

* Fixed custom domain hook so that it sets both Host header and URL correctly.

## 2.7.3
> Released on 2025/05/13

### Bug fixes

* Fixed JSON parsing errors when provisioning AWS VPC peering gateway in `konnect_cloud_gateway_transit_gateway` resource.

## 2.7.2
> Released on 2025/05/06

### Bug fixes

* Fixed output of `terraform plan` for `konnect_api_product_document` and `konnect_api_product_specification` resources.

## 2.7.1
> Released on 2025/05/01

### Bug fixes

* Fixed a bug in the `file_log` plugin `path` validation that disallowed `/dev/stdout` as a path

## 2.7.0  
> Released on 2025/05/01

### Features

* Add support for new Gateway 3.10 plugins:
  * `gateway_plugin_ai_rag_injector`
  * `gateway_plugin_ai_sanitizer`
  * `gateway_plugin_confluent_consume`
  * `gateway_plugin_kafka_consume`
  * `gateway_plugin_request_callout`

### Bug fixes

* The properties `limit` and `window_size` now accept array of values in `ai-rate-limiting-advanced` plugin according to the Gateway 3.10 schema

## 2.6.0
> Released on 2025/04/22

### Features

* Add support for `konnect_gateway_custom_plugin_streaming` resource, which allows custom plugins to be deployed to Dedicated Cloud Gateways

### Bug fixes

* The `scopes` property in `konnect_gateway_plugin_upstream_oauth` can now be set to an empty array

## 2.5.0
> Released on 2025/04/15

### Features

* Add support for `konnect_gateway_control_plane_list`  data source

### Bug fixes

* The `scopes` property in `konnect_gateway_plugin_openid_connect` is no longer required

## 2.4.1
> Released on 2025/03/13

### Bug fixes

* Fixed support for `headers` in the `konnect_gateway_route` resource
* Removed the `gateway_plugin_konnect_application_auth` plugin as it is a system-managed plugin and can not be configured by users

## 2.4.0
> Released on 2025/03/05

### Features

* Add support for `konnect_gateway_route_expression` resource
* `konnect_mesh_control_plane` now supports `features` property
* Add support for AWS VPC peering gateway in `konnect_cloud_gateway_transit_gateway` resource

### Bug fixes

* Foreign key references (e.g. `service.id`) in Kong Gateway are now cleared when removed from the manifest

## 2.3.0
> Released on 2025/02/10

### Features

* Add support for `konnect_portal_team` resource
* Add support for `konnect_audit_log` resource
* Add support for `konnect_audit_log_destination` resource
* Add support for `konnect_gateway_config_store` resource
* Add support for new Gateway 3.9 plugins:
  * `ai-azure-content-safety`
  * `confluent`
  * `service-protection`
  * `standard-webhooks`
  * `ai-semantic-cache`
  * `upstream-oauth`
  * `injection-protection`
  * `redirect`
  * `datadog-tracing`
  * `ai-proxy-advanced`
  * `ai-semantic-prompt-guard`
  * `json-threat-protection`
  * `header-cert-auth`
  * `ai-rate-limiting-advanced`

### Bug fixes

* Make `konnect_cloud_gateway_network` plan output deterministic
* Move `protocols` and foreign key (e.g. `service`, `route`) validation for plugins to be at plan time rather than runtime.

## 2.2.0
> Released on 2024/12/06

### Features

* Add support for Portal SAML authentication
* Added support for Konnect India region

## 2.1.0
> Released on 2024/12/06

### Features

* Add support for `konnect_gateway_mtls_auth` resource
* API Products now support `public_labels`
* Added support for Konnect Middle-East region

### Bug Fixes

* Fix `Expected #sdk.Konnect, got: *sdk.Konnect.` error with custom plugins (and add an integration test)

## 2.0.2
> Released on 2024/11/12

### Bug Fixes

* Removing service, route, consumer etc scope from a plugin now sends `null` to the API, removing the link between the plugin and the resource.

## 2.0.1
> Released on 2024/11/11

### Bug Fixes

* Switched `CLUSTER_TYPE_HYBRID` to a deprecation rather than a hard removal as changing the `type` results in resource replacement

## 2.0.0
> Released on 2024/11/11

### 🚨 BREAKING CHANGES
* The `konnect_cloud_gateway_transit_gateway` resource structure has had a bug fix that enables both AWS and Azure support. This has changed the resource structure and you will need to recreate any transit gateway attachments.
* `CLUSTER_TYPE_HYBRID` is an internal cluster name which has now been removed from the provider. Use `CLUSTER_TYPE_CONTROL_PLANE` instead. 
* Setting `scopes = []` in the OpenID Connect plugin is now respected. This means that the default has changed from `['openid']` to `[]`.

### Features

* Updated Kong Gateway plugin schemas to 3.8.x
* Added support for Azure Cloud Gateways deployments

### Bug Fixes

* Fix OpenAPI filename validation in API Products when using variables

## 1.0.0
> Released on 2024/09/11

### Features

* Released `terraform-provider-konnect` as GA 🎉

## 0.7.0
> Released on 2024/09/11

### Features
* Add the `konnect_gateway_custom_plugin` resource for managing non-bundled plugins
* Add support for all bundled Kong plugins
* Add support for `labels` to `konnect_api_product_version`
* Add support for `ordering` to `konnect_gateway_plugin`

### Bug fixes
* Force recreation of `konnect_gateway_*` entities if `control_plane_id` changes
* `proxy_urls` on `konnect_control_plane` and `dataplane_groups` on `konnect_cloud_gateway` are now marked as sets rather than list. This makes `terraform diff` ignore the order of items in the list.

## 0.6.3
> Released on 2024/08/25

### Features
* Add support for `labels` to `application_auth_strategies`

### Bug fixes
* The `cluster_cert_key` in `serverless_cloud_gateway` is now marked as `sensitive`

## 0.6.2
> Released on 2024/07/24

### Features
* Add the ability for the Kong team to use the provider against development environments

### Bug fixes
* Generate up to date documentation

## 0.6.1
> Released on 2024/07/23

### Features
* Add list of `enums` for available roles when using the Identity API
* Add pattern validation based on OpenAPI spec when running `terraform validate`

### Bug Fixes
* Fixed `terraform import` for the `konnect_portal` resource

## 0.6.0
> Released on 2024/07/11

### BREAKING CHANGES
* The `konnect_portal` resource can no longer be used to adopt the existing default portal without first running `terraform import`

### Features
* Add support for the  `konnect_api_product_document`, `konnect_portal_appearance`,  `konnect_gateway_plugin_pre_function`, `konnect_gateway_plugin_post_function` and `konnect_gateway_plugin_statsd` resources

## 0.5.1
> Released on 2024/07/03

### Features
* Add `tfdocs` for all existing resources

## 0.5.0
> Released on 2024/07/03

### Features
* Add support for the `konnect_serverless_cloud_gateway`, `konnect_gateway_plugin_exit_transformer`, `konnect_gateway_consumer_group` and `konnect_gateway_consumer_group_member` resources

### Bug Fixes
* `konnect_gateway_*` resources are now updated in place rather than any change causing the entity to be recreated

## 0.4.0
> Released on 2024/06/24

### Features
* Add support for the `konnect_api_product_specification`, `konnect_gateway_custom_plugin_schema`, `konnect_team`, `konnect_team_role` and `konnect_team_user` resources

### Bug Fixes
* Make Cloud Gateways + Identity APIs point at `global.api.konghq.com` rather than the provided `server_url`

## 0.3.1
> Released on 2024/06/14

### Bug Fixes
* Ensure that `konnect_gateway_vault.config` is sent via the provider

## 0.3.0
> Released on 2024/06/14

### Features
* Update all schemas to Kong Gateway 3.7.x

### Bug Fixes
* Do not send default values from the provider as they'll be applied on the server automatically
  * This allows terraform-provider-konnect 0.3.0 to be used with Kong Gateway 3.6 and below

## 0.2.4
> Released on 2024/05/31

### Features
* Allow users to set provider attributes via environment variables
  * `personal_access_token` = `KONNECT_TOKEN`
  * `system_account_access_token` = `KONNECT_SPAT`
  * `server_url` = `KONNECT_SERVER_URL`

## 0.2.3
> Released on 2024/05/14

### Features
* Add support for the `konnect_cloud_gateway_configuration`, `konnect_cloud_gateway_custom_domain`, `konnect_cloud_gateway_network` and `konnect_cloud_gateway_transit_gateway` resources

### Bug Fixes
* Fix `konnect_portal_product_version` creation bug

## 0.2.2
> Released on 2024/05/04

### Features
* Add support for the `konnect_system_account`, `konnect_system_account_access_token`, `konnect_system_account_role` and `konnect_system_account_team` resources

## 0.2.1
> Released on 2024/04/26

### Features
* Add support for the `konnect_gateway_data_plane_client_certificate` resource

## 0.2.0
> Released on 2024/04/18

### BREAKING CHANGES
* The provider `server_url` no longer contains an API version. Update your provider configuration to use `https://us.api.konghq.com` (or your chosen region)
* `labels` on the `konnect_gateway_control_plane` resource now accept a map of values rather than using `jsonencode()`

### Fixes
* Mesh control planes can be created in regions other than NA
* Gateway control planes now accept `CLUSTER_TYPE_CONTROL_PLANE` in the `cluster_type` field

## 0.1.0
> Released on 2024/04/16

* Initial release
