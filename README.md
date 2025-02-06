undefined
<!-- Start Summary [summary] -->
## Summary

Konnect API: The Konnect platform API

For more information about the API: [Documentation for Kong Gateway and its APIs](https://docs.konghq.com)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
  * [Installation](#installation)
  * [Available Resources and Data Sources](#available-resources-and-data-sources)
  * [Testing the provider locally](#testing-the-provider-locally)

<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "2.2.0"
    }
  }
}

provider "konnect" {
  # Configuration options
}
```
<!-- End Installation [installation] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [konnect_api_product](docs/resources/api_product.md)
* [konnect_api_product_document](docs/resources/api_product_document.md)
* [konnect_api_product_specification](docs/resources/api_product_specification.md)
* [konnect_api_product_version](docs/resources/api_product_version.md)
* [konnect_application_auth_strategy](docs/resources/application_auth_strategy.md)
* [konnect_audit_log](docs/resources/audit_log.md)
* [konnect_cloud_gateway_configuration](docs/resources/cloud_gateway_configuration.md)
* [konnect_cloud_gateway_custom_domain](docs/resources/cloud_gateway_custom_domain.md)
* [konnect_cloud_gateway_network](docs/resources/cloud_gateway_network.md)
* [konnect_cloud_gateway_transit_gateway](docs/resources/cloud_gateway_transit_gateway.md)
* [konnect_gateway_acl](docs/resources/gateway_acl.md)
* [konnect_gateway_basic_auth](docs/resources/gateway_basic_auth.md)
* [konnect_gateway_ca_certificate](docs/resources/gateway_ca_certificate.md)
* [konnect_gateway_certificate](docs/resources/gateway_certificate.md)
* [konnect_gateway_consumer](docs/resources/gateway_consumer.md)
* [konnect_gateway_consumer_group](docs/resources/gateway_consumer_group.md)
* [konnect_gateway_consumer_group_member](docs/resources/gateway_consumer_group_member.md)
* [konnect_gateway_control_plane](docs/resources/gateway_control_plane.md)
* [konnect_gateway_control_plane_membership](docs/resources/gateway_control_plane_membership.md)
* [konnect_gateway_custom_plugin_schema](docs/resources/gateway_custom_plugin_schema.md)
* [konnect_gateway_data_plane_client_certificate](docs/resources/gateway_data_plane_client_certificate.md)
* [konnect_gateway_hmac_auth](docs/resources/gateway_hmac_auth.md)
* [konnect_gateway_jwt](docs/resources/gateway_jwt.md)
* [konnect_gateway_key](docs/resources/gateway_key.md)
* [konnect_gateway_key_auth](docs/resources/gateway_key_auth.md)
* [konnect_gateway_key_set](docs/resources/gateway_key_set.md)
* [konnect_gateway_mtls_auth](docs/resources/gateway_mtls_auth.md)
* [konnect_gateway_plugin_acl](docs/resources/gateway_plugin_acl.md)
* [konnect_gateway_plugin_acme](docs/resources/gateway_plugin_acme.md)
* [konnect_gateway_plugin_ai_azure_content_safety](docs/resources/gateway_plugin_ai_azure_content_safety.md)
* [konnect_gateway_plugin_ai_prompt_decorator](docs/resources/gateway_plugin_ai_prompt_decorator.md)
* [konnect_gateway_plugin_ai_prompt_guard](docs/resources/gateway_plugin_ai_prompt_guard.md)
* [konnect_gateway_plugin_ai_prompt_template](docs/resources/gateway_plugin_ai_prompt_template.md)
* [konnect_gateway_plugin_ai_proxy](docs/resources/gateway_plugin_ai_proxy.md)
* [konnect_gateway_plugin_ai_proxy_advanced](docs/resources/gateway_plugin_ai_proxy_advanced.md)
* [konnect_gateway_plugin_ai_rate_limiting_advanced](docs/resources/gateway_plugin_ai_rate_limiting_advanced.md)
* [konnect_gateway_plugin_ai_request_transformer](docs/resources/gateway_plugin_ai_request_transformer.md)
* [konnect_gateway_plugin_ai_response_transformer](docs/resources/gateway_plugin_ai_response_transformer.md)
* [konnect_gateway_plugin_ai_semantic_cache](docs/resources/gateway_plugin_ai_semantic_cache.md)
* [konnect_gateway_plugin_ai_semantic_prompt_guard](docs/resources/gateway_plugin_ai_semantic_prompt_guard.md)
* [konnect_gateway_plugin_aws_lambda](docs/resources/gateway_plugin_aws_lambda.md)
* [konnect_gateway_plugin_azure_functions](docs/resources/gateway_plugin_azure_functions.md)
* [konnect_gateway_plugin_basic_auth](docs/resources/gateway_plugin_basic_auth.md)
* [konnect_gateway_plugin_bot_detection](docs/resources/gateway_plugin_bot_detection.md)
* [konnect_gateway_plugin_canary](docs/resources/gateway_plugin_canary.md)
* [konnect_gateway_plugin_confluent](docs/resources/gateway_plugin_confluent.md)
* [konnect_gateway_plugin_correlation_id](docs/resources/gateway_plugin_correlation_id.md)
* [konnect_gateway_plugin_cors](docs/resources/gateway_plugin_cors.md)
* [konnect_gateway_plugin_datadog](docs/resources/gateway_plugin_datadog.md)
* [konnect_gateway_plugin_datadog_tracing](docs/resources/gateway_plugin_datadog_tracing.md)
* [konnect_gateway_plugin_degraphql](docs/resources/gateway_plugin_degraphql.md)
* [konnect_gateway_plugin_exit_transformer](docs/resources/gateway_plugin_exit_transformer.md)
* [konnect_gateway_plugin_file_log](docs/resources/gateway_plugin_file_log.md)
* [konnect_gateway_plugin_forward_proxy](docs/resources/gateway_plugin_forward_proxy.md)
* [konnect_gateway_plugin_graphql_proxy_cache_advanced](docs/resources/gateway_plugin_graphql_proxy_cache_advanced.md)
* [konnect_gateway_plugin_graphql_rate_limiting_advanced](docs/resources/gateway_plugin_graphql_rate_limiting_advanced.md)
* [konnect_gateway_plugin_grpc_gateway](docs/resources/gateway_plugin_grpc_gateway.md)
* [konnect_gateway_plugin_grpc_web](docs/resources/gateway_plugin_grpc_web.md)
* [konnect_gateway_plugin_header_cert_auth](docs/resources/gateway_plugin_header_cert_auth.md)
* [konnect_gateway_plugin_hmac_auth](docs/resources/gateway_plugin_hmac_auth.md)
* [konnect_gateway_plugin_http_log](docs/resources/gateway_plugin_http_log.md)
* [konnect_gateway_plugin_injection_protection](docs/resources/gateway_plugin_injection_protection.md)
* [konnect_gateway_plugin_ip_restriction](docs/resources/gateway_plugin_ip_restriction.md)
* [konnect_gateway_plugin_jq](docs/resources/gateway_plugin_jq.md)
* [konnect_gateway_plugin_json_threat_protection](docs/resources/gateway_plugin_json_threat_protection.md)
* [konnect_gateway_plugin_jwe_decrypt](docs/resources/gateway_plugin_jwe_decrypt.md)
* [konnect_gateway_plugin_jwt](docs/resources/gateway_plugin_jwt.md)
* [konnect_gateway_plugin_jwt_signer](docs/resources/gateway_plugin_jwt_signer.md)
* [konnect_gateway_plugin_kafka_log](docs/resources/gateway_plugin_kafka_log.md)
* [konnect_gateway_plugin_kafka_upstream](docs/resources/gateway_plugin_kafka_upstream.md)
* [konnect_gateway_plugin_key_auth](docs/resources/gateway_plugin_key_auth.md)
* [konnect_gateway_plugin_key_auth_enc](docs/resources/gateway_plugin_key_auth_enc.md)
* [konnect_gateway_plugin_konnect_application_auth](docs/resources/gateway_plugin_konnect_application_auth.md)
* [konnect_gateway_plugin_ldap_auth](docs/resources/gateway_plugin_ldap_auth.md)
* [konnect_gateway_plugin_ldap_auth_advanced](docs/resources/gateway_plugin_ldap_auth_advanced.md)
* [konnect_gateway_plugin_loggly](docs/resources/gateway_plugin_loggly.md)
* [konnect_gateway_plugin_mocking](docs/resources/gateway_plugin_mocking.md)
* [konnect_gateway_plugin_mtls_auth](docs/resources/gateway_plugin_mtls_auth.md)
* [konnect_gateway_plugin_oas_validation](docs/resources/gateway_plugin_oas_validation.md)
* [konnect_gateway_plugin_oauth2](docs/resources/gateway_plugin_oauth2.md)
* [konnect_gateway_plugin_oauth2_introspection](docs/resources/gateway_plugin_oauth2_introspection.md)
* [konnect_gateway_plugin_opa](docs/resources/gateway_plugin_opa.md)
* [konnect_gateway_plugin_openid_connect](docs/resources/gateway_plugin_openid_connect.md)
* [konnect_gateway_plugin_opentelemetry](docs/resources/gateway_plugin_opentelemetry.md)
* [konnect_gateway_plugin_post_function](docs/resources/gateway_plugin_post_function.md)
* [konnect_gateway_plugin_pre_function](docs/resources/gateway_plugin_pre_function.md)
* [konnect_gateway_plugin_prometheus](docs/resources/gateway_plugin_prometheus.md)
* [konnect_gateway_plugin_proxy_cache](docs/resources/gateway_plugin_proxy_cache.md)
* [konnect_gateway_plugin_proxy_cache_advanced](docs/resources/gateway_plugin_proxy_cache_advanced.md)
* [konnect_gateway_plugin_rate_limiting](docs/resources/gateway_plugin_rate_limiting.md)
* [konnect_gateway_plugin_rate_limiting_advanced](docs/resources/gateway_plugin_rate_limiting_advanced.md)
* [konnect_gateway_plugin_redirect](docs/resources/gateway_plugin_redirect.md)
* [konnect_gateway_plugin_request_size_limiting](docs/resources/gateway_plugin_request_size_limiting.md)
* [konnect_gateway_plugin_request_termination](docs/resources/gateway_plugin_request_termination.md)
* [konnect_gateway_plugin_request_transformer](docs/resources/gateway_plugin_request_transformer.md)
* [konnect_gateway_plugin_request_transformer_advanced](docs/resources/gateway_plugin_request_transformer_advanced.md)
* [konnect_gateway_plugin_request_validator](docs/resources/gateway_plugin_request_validator.md)
* [konnect_gateway_plugin_response_ratelimiting](docs/resources/gateway_plugin_response_ratelimiting.md)
* [konnect_gateway_plugin_response_transformer](docs/resources/gateway_plugin_response_transformer.md)
* [konnect_gateway_plugin_response_transformer_advanced](docs/resources/gateway_plugin_response_transformer_advanced.md)
* [konnect_gateway_plugin_route_by_header](docs/resources/gateway_plugin_route_by_header.md)
* [konnect_gateway_plugin_route_transformer_advanced](docs/resources/gateway_plugin_route_transformer_advanced.md)
* [konnect_gateway_plugin_saml](docs/resources/gateway_plugin_saml.md)
* [konnect_gateway_plugin_service_protection](docs/resources/gateway_plugin_service_protection.md)
* [konnect_gateway_plugin_session](docs/resources/gateway_plugin_session.md)
* [konnect_gateway_plugin_standard_webhooks](docs/resources/gateway_plugin_standard_webhooks.md)
* [konnect_gateway_plugin_statsd](docs/resources/gateway_plugin_statsd.md)
* [konnect_gateway_plugin_statsd_advanced](docs/resources/gateway_plugin_statsd_advanced.md)
* [konnect_gateway_plugin_syslog](docs/resources/gateway_plugin_syslog.md)
* [konnect_gateway_plugin_tcp_log](docs/resources/gateway_plugin_tcp_log.md)
* [konnect_gateway_plugin_tls_handshake_modifier](docs/resources/gateway_plugin_tls_handshake_modifier.md)
* [konnect_gateway_plugin_tls_metadata_headers](docs/resources/gateway_plugin_tls_metadata_headers.md)
* [konnect_gateway_plugin_udp_log](docs/resources/gateway_plugin_udp_log.md)
* [konnect_gateway_plugin_upstream_oauth](docs/resources/gateway_plugin_upstream_oauth.md)
* [konnect_gateway_plugin_upstream_timeout](docs/resources/gateway_plugin_upstream_timeout.md)
* [konnect_gateway_plugin_vault_auth](docs/resources/gateway_plugin_vault_auth.md)
* [konnect_gateway_plugin_websocket_size_limit](docs/resources/gateway_plugin_websocket_size_limit.md)
* [konnect_gateway_plugin_websocket_validator](docs/resources/gateway_plugin_websocket_validator.md)
* [konnect_gateway_plugin_xml_threat_protection](docs/resources/gateway_plugin_xml_threat_protection.md)
* [konnect_gateway_plugin_zipkin](docs/resources/gateway_plugin_zipkin.md)
* [konnect_gateway_route](docs/resources/gateway_route.md)
* [konnect_gateway_service](docs/resources/gateway_service.md)
* [konnect_gateway_sni](docs/resources/gateway_sni.md)
* [konnect_gateway_target](docs/resources/gateway_target.md)
* [konnect_gateway_upstream](docs/resources/gateway_upstream.md)
* [konnect_gateway_vault](docs/resources/gateway_vault.md)
* [konnect_hostname_generator](docs/resources/hostname_generator.md)
* [konnect_mesh](docs/resources/mesh.md)
* [konnect_mesh_access_log](docs/resources/mesh_access_log.md)
* [konnect_mesh_circuit_breaker](docs/resources/mesh_circuit_breaker.md)
* [konnect_mesh_control_plane](docs/resources/mesh_control_plane.md)
* [konnect_mesh_external_service](docs/resources/mesh_external_service.md)
* [konnect_mesh_fault_injection](docs/resources/mesh_fault_injection.md)
* [konnect_mesh_gateway](docs/resources/mesh_gateway.md)
* [konnect_mesh_global_rate_limit](docs/resources/mesh_global_rate_limit.md)
* [konnect_mesh_health_check](docs/resources/mesh_health_check.md)
* [konnect_mesh_http_route](docs/resources/mesh_http_route.md)
* [konnect_mesh_load_balancing_strategy](docs/resources/mesh_load_balancing_strategy.md)
* [konnect_mesh_metric](docs/resources/mesh_metric.md)
* [konnect_mesh_multi_zone_service](docs/resources/mesh_multi_zone_service.md)
* [konnect_mesh_opa](docs/resources/mesh_opa.md)
* [konnect_mesh_passthrough](docs/resources/mesh_passthrough.md)
* [konnect_mesh_proxy_patch](docs/resources/mesh_proxy_patch.md)
* [konnect_mesh_rate_limit](docs/resources/mesh_rate_limit.md)
* [konnect_mesh_retry](docs/resources/mesh_retry.md)
* [konnect_mesh_service](docs/resources/mesh_service.md)
* [konnect_mesh_tcp_route](docs/resources/mesh_tcp_route.md)
* [konnect_mesh_timeout](docs/resources/mesh_timeout.md)
* [konnect_mesh_tls](docs/resources/mesh_tls.md)
* [konnect_mesh_trace](docs/resources/mesh_trace.md)
* [konnect_mesh_traffic_permission](docs/resources/mesh_traffic_permission.md)
* [konnect_portal](docs/resources/portal.md)
* [konnect_portal_appearance](docs/resources/portal_appearance.md)
* [konnect_portal_auth](docs/resources/portal_auth.md)
* [konnect_portal_product_version](docs/resources/portal_product_version.md)
* [konnect_portal_team](docs/resources/portal_team.md)
* [konnect_serverless_cloud_gateway](docs/resources/serverless_cloud_gateway.md)
* [konnect_system_account](docs/resources/system_account.md)
* [konnect_system_account_access_token](docs/resources/system_account_access_token.md)
* [konnect_system_account_role](docs/resources/system_account_role.md)
* [konnect_system_account_team](docs/resources/system_account_team.md)
* [konnect_team](docs/resources/team.md)
* [konnect_team_role](docs/resources/team_role.md)
* [konnect_team_user](docs/resources/team_user.md)
### Data Sources

* [konnect_api_product](docs/data-sources/api_product.md)
* [konnect_api_product_document](docs/data-sources/api_product_document.md)
* [konnect_api_product_specification](docs/data-sources/api_product_specification.md)
* [konnect_api_product_version](docs/data-sources/api_product_version.md)
* [konnect_application_auth_strategy](docs/data-sources/application_auth_strategy.md)
* [konnect_audit_log](docs/data-sources/audit_log.md)
* [konnect_cloud_gateway_configuration](docs/data-sources/cloud_gateway_configuration.md)
* [konnect_cloud_gateway_custom_domain](docs/data-sources/cloud_gateway_custom_domain.md)
* [konnect_cloud_gateway_network](docs/data-sources/cloud_gateway_network.md)
* [konnect_cloud_gateway_provider_account_list](docs/data-sources/cloud_gateway_provider_account_list.md)
* [konnect_cloud_gateway_transit_gateway](docs/data-sources/cloud_gateway_transit_gateway.md)
* [konnect_gateway_acl](docs/data-sources/gateway_acl.md)
* [konnect_gateway_basic_auth](docs/data-sources/gateway_basic_auth.md)
* [konnect_gateway_ca_certificate](docs/data-sources/gateway_ca_certificate.md)
* [konnect_gateway_certificate](docs/data-sources/gateway_certificate.md)
* [konnect_gateway_consumer](docs/data-sources/gateway_consumer.md)
* [konnect_gateway_consumer_group](docs/data-sources/gateway_consumer_group.md)
* [konnect_gateway_control_plane](docs/data-sources/gateway_control_plane.md)
* [konnect_gateway_custom_plugin_schema](docs/data-sources/gateway_custom_plugin_schema.md)
* [konnect_gateway_data_plane_client_certificate](docs/data-sources/gateway_data_plane_client_certificate.md)
* [konnect_gateway_hmac_auth](docs/data-sources/gateway_hmac_auth.md)
* [konnect_gateway_jwt](docs/data-sources/gateway_jwt.md)
* [konnect_gateway_key](docs/data-sources/gateway_key.md)
* [konnect_gateway_key_auth](docs/data-sources/gateway_key_auth.md)
* [konnect_gateway_key_set](docs/data-sources/gateway_key_set.md)
* [konnect_gateway_mtls_auth](docs/data-sources/gateway_mtls_auth.md)
* [konnect_gateway_plugin_acl](docs/data-sources/gateway_plugin_acl.md)
* [konnect_gateway_plugin_acme](docs/data-sources/gateway_plugin_acme.md)
* [konnect_gateway_plugin_ai_azure_content_safety](docs/data-sources/gateway_plugin_ai_azure_content_safety.md)
* [konnect_gateway_plugin_ai_prompt_decorator](docs/data-sources/gateway_plugin_ai_prompt_decorator.md)
* [konnect_gateway_plugin_ai_prompt_guard](docs/data-sources/gateway_plugin_ai_prompt_guard.md)
* [konnect_gateway_plugin_ai_prompt_template](docs/data-sources/gateway_plugin_ai_prompt_template.md)
* [konnect_gateway_plugin_ai_proxy](docs/data-sources/gateway_plugin_ai_proxy.md)
* [konnect_gateway_plugin_ai_proxy_advanced](docs/data-sources/gateway_plugin_ai_proxy_advanced.md)
* [konnect_gateway_plugin_ai_rate_limiting_advanced](docs/data-sources/gateway_plugin_ai_rate_limiting_advanced.md)
* [konnect_gateway_plugin_ai_request_transformer](docs/data-sources/gateway_plugin_ai_request_transformer.md)
* [konnect_gateway_plugin_ai_response_transformer](docs/data-sources/gateway_plugin_ai_response_transformer.md)
* [konnect_gateway_plugin_ai_semantic_cache](docs/data-sources/gateway_plugin_ai_semantic_cache.md)
* [konnect_gateway_plugin_ai_semantic_prompt_guard](docs/data-sources/gateway_plugin_ai_semantic_prompt_guard.md)
* [konnect_gateway_plugin_aws_lambda](docs/data-sources/gateway_plugin_aws_lambda.md)
* [konnect_gateway_plugin_azure_functions](docs/data-sources/gateway_plugin_azure_functions.md)
* [konnect_gateway_plugin_basic_auth](docs/data-sources/gateway_plugin_basic_auth.md)
* [konnect_gateway_plugin_bot_detection](docs/data-sources/gateway_plugin_bot_detection.md)
* [konnect_gateway_plugin_canary](docs/data-sources/gateway_plugin_canary.md)
* [konnect_gateway_plugin_confluent](docs/data-sources/gateway_plugin_confluent.md)
* [konnect_gateway_plugin_correlation_id](docs/data-sources/gateway_plugin_correlation_id.md)
* [konnect_gateway_plugin_cors](docs/data-sources/gateway_plugin_cors.md)
* [konnect_gateway_plugin_datadog](docs/data-sources/gateway_plugin_datadog.md)
* [konnect_gateway_plugin_datadog_tracing](docs/data-sources/gateway_plugin_datadog_tracing.md)
* [konnect_gateway_plugin_degraphql](docs/data-sources/gateway_plugin_degraphql.md)
* [konnect_gateway_plugin_exit_transformer](docs/data-sources/gateway_plugin_exit_transformer.md)
* [konnect_gateway_plugin_file_log](docs/data-sources/gateway_plugin_file_log.md)
* [konnect_gateway_plugin_forward_proxy](docs/data-sources/gateway_plugin_forward_proxy.md)
* [konnect_gateway_plugin_graphql_proxy_cache_advanced](docs/data-sources/gateway_plugin_graphql_proxy_cache_advanced.md)
* [konnect_gateway_plugin_graphql_rate_limiting_advanced](docs/data-sources/gateway_plugin_graphql_rate_limiting_advanced.md)
* [konnect_gateway_plugin_grpc_gateway](docs/data-sources/gateway_plugin_grpc_gateway.md)
* [konnect_gateway_plugin_grpc_web](docs/data-sources/gateway_plugin_grpc_web.md)
* [konnect_gateway_plugin_header_cert_auth](docs/data-sources/gateway_plugin_header_cert_auth.md)
* [konnect_gateway_plugin_hmac_auth](docs/data-sources/gateway_plugin_hmac_auth.md)
* [konnect_gateway_plugin_http_log](docs/data-sources/gateway_plugin_http_log.md)
* [konnect_gateway_plugin_injection_protection](docs/data-sources/gateway_plugin_injection_protection.md)
* [konnect_gateway_plugin_ip_restriction](docs/data-sources/gateway_plugin_ip_restriction.md)
* [konnect_gateway_plugin_jq](docs/data-sources/gateway_plugin_jq.md)
* [konnect_gateway_plugin_json_threat_protection](docs/data-sources/gateway_plugin_json_threat_protection.md)
* [konnect_gateway_plugin_jwe_decrypt](docs/data-sources/gateway_plugin_jwe_decrypt.md)
* [konnect_gateway_plugin_jwt](docs/data-sources/gateway_plugin_jwt.md)
* [konnect_gateway_plugin_jwt_signer](docs/data-sources/gateway_plugin_jwt_signer.md)
* [konnect_gateway_plugin_kafka_log](docs/data-sources/gateway_plugin_kafka_log.md)
* [konnect_gateway_plugin_kafka_upstream](docs/data-sources/gateway_plugin_kafka_upstream.md)
* [konnect_gateway_plugin_key_auth](docs/data-sources/gateway_plugin_key_auth.md)
* [konnect_gateway_plugin_key_auth_enc](docs/data-sources/gateway_plugin_key_auth_enc.md)
* [konnect_gateway_plugin_konnect_application_auth](docs/data-sources/gateway_plugin_konnect_application_auth.md)
* [konnect_gateway_plugin_ldap_auth](docs/data-sources/gateway_plugin_ldap_auth.md)
* [konnect_gateway_plugin_ldap_auth_advanced](docs/data-sources/gateway_plugin_ldap_auth_advanced.md)
* [konnect_gateway_plugin_loggly](docs/data-sources/gateway_plugin_loggly.md)
* [konnect_gateway_plugin_mocking](docs/data-sources/gateway_plugin_mocking.md)
* [konnect_gateway_plugin_mtls_auth](docs/data-sources/gateway_plugin_mtls_auth.md)
* [konnect_gateway_plugin_oas_validation](docs/data-sources/gateway_plugin_oas_validation.md)
* [konnect_gateway_plugin_oauth2](docs/data-sources/gateway_plugin_oauth2.md)
* [konnect_gateway_plugin_oauth2_introspection](docs/data-sources/gateway_plugin_oauth2_introspection.md)
* [konnect_gateway_plugin_opa](docs/data-sources/gateway_plugin_opa.md)
* [konnect_gateway_plugin_openid_connect](docs/data-sources/gateway_plugin_openid_connect.md)
* [konnect_gateway_plugin_opentelemetry](docs/data-sources/gateway_plugin_opentelemetry.md)
* [konnect_gateway_plugin_post_function](docs/data-sources/gateway_plugin_post_function.md)
* [konnect_gateway_plugin_pre_function](docs/data-sources/gateway_plugin_pre_function.md)
* [konnect_gateway_plugin_prometheus](docs/data-sources/gateway_plugin_prometheus.md)
* [konnect_gateway_plugin_proxy_cache](docs/data-sources/gateway_plugin_proxy_cache.md)
* [konnect_gateway_plugin_proxy_cache_advanced](docs/data-sources/gateway_plugin_proxy_cache_advanced.md)
* [konnect_gateway_plugin_rate_limiting](docs/data-sources/gateway_plugin_rate_limiting.md)
* [konnect_gateway_plugin_rate_limiting_advanced](docs/data-sources/gateway_plugin_rate_limiting_advanced.md)
* [konnect_gateway_plugin_redirect](docs/data-sources/gateway_plugin_redirect.md)
* [konnect_gateway_plugin_request_size_limiting](docs/data-sources/gateway_plugin_request_size_limiting.md)
* [konnect_gateway_plugin_request_termination](docs/data-sources/gateway_plugin_request_termination.md)
* [konnect_gateway_plugin_request_transformer](docs/data-sources/gateway_plugin_request_transformer.md)
* [konnect_gateway_plugin_request_transformer_advanced](docs/data-sources/gateway_plugin_request_transformer_advanced.md)
* [konnect_gateway_plugin_request_validator](docs/data-sources/gateway_plugin_request_validator.md)
* [konnect_gateway_plugin_response_ratelimiting](docs/data-sources/gateway_plugin_response_ratelimiting.md)
* [konnect_gateway_plugin_response_transformer](docs/data-sources/gateway_plugin_response_transformer.md)
* [konnect_gateway_plugin_response_transformer_advanced](docs/data-sources/gateway_plugin_response_transformer_advanced.md)
* [konnect_gateway_plugin_route_by_header](docs/data-sources/gateway_plugin_route_by_header.md)
* [konnect_gateway_plugin_route_transformer_advanced](docs/data-sources/gateway_plugin_route_transformer_advanced.md)
* [konnect_gateway_plugin_saml](docs/data-sources/gateway_plugin_saml.md)
* [konnect_gateway_plugin_service_protection](docs/data-sources/gateway_plugin_service_protection.md)
* [konnect_gateway_plugin_session](docs/data-sources/gateway_plugin_session.md)
* [konnect_gateway_plugin_standard_webhooks](docs/data-sources/gateway_plugin_standard_webhooks.md)
* [konnect_gateway_plugin_statsd](docs/data-sources/gateway_plugin_statsd.md)
* [konnect_gateway_plugin_statsd_advanced](docs/data-sources/gateway_plugin_statsd_advanced.md)
* [konnect_gateway_plugin_syslog](docs/data-sources/gateway_plugin_syslog.md)
* [konnect_gateway_plugin_tcp_log](docs/data-sources/gateway_plugin_tcp_log.md)
* [konnect_gateway_plugin_tls_handshake_modifier](docs/data-sources/gateway_plugin_tls_handshake_modifier.md)
* [konnect_gateway_plugin_tls_metadata_headers](docs/data-sources/gateway_plugin_tls_metadata_headers.md)
* [konnect_gateway_plugin_udp_log](docs/data-sources/gateway_plugin_udp_log.md)
* [konnect_gateway_plugin_upstream_oauth](docs/data-sources/gateway_plugin_upstream_oauth.md)
* [konnect_gateway_plugin_upstream_timeout](docs/data-sources/gateway_plugin_upstream_timeout.md)
* [konnect_gateway_plugin_vault_auth](docs/data-sources/gateway_plugin_vault_auth.md)
* [konnect_gateway_plugin_websocket_size_limit](docs/data-sources/gateway_plugin_websocket_size_limit.md)
* [konnect_gateway_plugin_websocket_validator](docs/data-sources/gateway_plugin_websocket_validator.md)
* [konnect_gateway_plugin_xml_threat_protection](docs/data-sources/gateway_plugin_xml_threat_protection.md)
* [konnect_gateway_plugin_zipkin](docs/data-sources/gateway_plugin_zipkin.md)
* [konnect_gateway_route](docs/data-sources/gateway_route.md)
* [konnect_gateway_service](docs/data-sources/gateway_service.md)
* [konnect_gateway_sni](docs/data-sources/gateway_sni.md)
* [konnect_gateway_target](docs/data-sources/gateway_target.md)
* [konnect_gateway_upstream](docs/data-sources/gateway_upstream.md)
* [konnect_gateway_vault](docs/data-sources/gateway_vault.md)
* [konnect_hostname_generator](docs/data-sources/hostname_generator.md)
* [konnect_hostname_generator_list](docs/data-sources/hostname_generator_list.md)
* [konnect_mesh](docs/data-sources/mesh.md)
* [konnect_mesh_access_log](docs/data-sources/mesh_access_log.md)
* [konnect_mesh_access_log_list](docs/data-sources/mesh_access_log_list.md)
* [konnect_mesh_circuit_breaker](docs/data-sources/mesh_circuit_breaker.md)
* [konnect_mesh_circuit_breaker_list](docs/data-sources/mesh_circuit_breaker_list.md)
* [konnect_mesh_control_plane](docs/data-sources/mesh_control_plane.md)
* [konnect_mesh_external_service](docs/data-sources/mesh_external_service.md)
* [konnect_mesh_external_service_list](docs/data-sources/mesh_external_service_list.md)
* [konnect_mesh_fault_injection](docs/data-sources/mesh_fault_injection.md)
* [konnect_mesh_fault_injection_list](docs/data-sources/mesh_fault_injection_list.md)
* [konnect_mesh_gateway](docs/data-sources/mesh_gateway.md)
* [konnect_mesh_gateway_list](docs/data-sources/mesh_gateway_list.md)
* [konnect_mesh_global_rate_limit](docs/data-sources/mesh_global_rate_limit.md)
* [konnect_mesh_global_rate_limit_list](docs/data-sources/mesh_global_rate_limit_list.md)
* [konnect_mesh_health_check](docs/data-sources/mesh_health_check.md)
* [konnect_mesh_health_check_list](docs/data-sources/mesh_health_check_list.md)
* [konnect_mesh_http_route](docs/data-sources/mesh_http_route.md)
* [konnect_mesh_http_route_list](docs/data-sources/mesh_http_route_list.md)
* [konnect_mesh_list](docs/data-sources/mesh_list.md)
* [konnect_mesh_load_balancing_strategy](docs/data-sources/mesh_load_balancing_strategy.md)
* [konnect_mesh_load_balancing_strategy_list](docs/data-sources/mesh_load_balancing_strategy_list.md)
* [konnect_mesh_metric](docs/data-sources/mesh_metric.md)
* [konnect_mesh_metric_list](docs/data-sources/mesh_metric_list.md)
* [konnect_mesh_multi_zone_service](docs/data-sources/mesh_multi_zone_service.md)
* [konnect_mesh_multi_zone_service_list](docs/data-sources/mesh_multi_zone_service_list.md)
* [konnect_mesh_opa](docs/data-sources/mesh_opa.md)
* [konnect_mesh_opa_list](docs/data-sources/mesh_opa_list.md)
* [konnect_mesh_passthrough](docs/data-sources/mesh_passthrough.md)
* [konnect_mesh_passthrough_list](docs/data-sources/mesh_passthrough_list.md)
* [konnect_mesh_proxy_patch](docs/data-sources/mesh_proxy_patch.md)
* [konnect_mesh_proxy_patch_list](docs/data-sources/mesh_proxy_patch_list.md)
* [konnect_mesh_rate_limit](docs/data-sources/mesh_rate_limit.md)
* [konnect_mesh_rate_limit_list](docs/data-sources/mesh_rate_limit_list.md)
* [konnect_mesh_retry](docs/data-sources/mesh_retry.md)
* [konnect_mesh_retry_list](docs/data-sources/mesh_retry_list.md)
* [konnect_mesh_service](docs/data-sources/mesh_service.md)
* [konnect_mesh_service_list](docs/data-sources/mesh_service_list.md)
* [konnect_mesh_tcp_route](docs/data-sources/mesh_tcp_route.md)
* [konnect_mesh_tcp_route_list](docs/data-sources/mesh_tcp_route_list.md)
* [konnect_mesh_timeout](docs/data-sources/mesh_timeout.md)
* [konnect_mesh_timeout_list](docs/data-sources/mesh_timeout_list.md)
* [konnect_mesh_tls](docs/data-sources/mesh_tls.md)
* [konnect_mesh_tls_list](docs/data-sources/mesh_tls_list.md)
* [konnect_mesh_trace](docs/data-sources/mesh_trace.md)
* [konnect_mesh_trace_list](docs/data-sources/mesh_trace_list.md)
* [konnect_mesh_traffic_permission](docs/data-sources/mesh_traffic_permission.md)
* [konnect_mesh_traffic_permission_list](docs/data-sources/mesh_traffic_permission_list.md)
* [konnect_portal_appearance](docs/data-sources/portal_appearance.md)
* [konnect_portal_auth](docs/data-sources/portal_auth.md)
* [konnect_portal_list](docs/data-sources/portal_list.md)
* [konnect_portal_product_version](docs/data-sources/portal_product_version.md)
* [konnect_portal_team](docs/data-sources/portal_team.md)
* [konnect_serverless_cloud_gateway](docs/data-sources/serverless_cloud_gateway.md)
* [konnect_system_account](docs/data-sources/system_account.md)
* [konnect_system_account_access_token](docs/data-sources/system_account_access_token.md)
* [konnect_team](docs/data-sources/team.md)
<!-- End Available Resources and Data Sources [operations] -->

<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-konnect`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/kong/konnect" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
