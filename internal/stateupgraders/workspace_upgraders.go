// Per-resource `workspace` backfill upgraders, consolidated into one file
// instead of Speakeasy's usual one-per-entity layout. See .genignore.
package stateupgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func GatewayaclStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_acl", req, resp, defaultWorkspaceIfMissing)
}

func GatewaybasicauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_basic_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaycacertificateStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_ca_certificate", req, resp, defaultWorkspaceIfMissing)
}

func GatewaycertificateStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_certificate", req, resp, defaultWorkspaceIfMissing)
}

func GatewayconsumerStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_consumer", req, resp, defaultWorkspaceIfMissing)
}

func GatewayconsumergroupStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_consumer_group", req, resp, defaultWorkspaceIfMissing)
}

func GatewayconsumergroupmemberStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_consumer_group_member", req, resp, defaultWorkspaceIfMissing)
}

func GatewayhmacauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_hmac_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewayjwtStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_jwt", req, resp, defaultWorkspaceIfMissing)
}

func GatewaykeyStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_key", req, resp, defaultWorkspaceIfMissing)
}

func GatewaykeyauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_key_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaykeysetStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_key_set", req, resp, defaultWorkspaceIfMissing)
}

func GatewaymtlsauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_mtls_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypartialStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_partial", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaceStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ace", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaclStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_acl", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginacmeStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_acme", req, resp, defaultWorkspaceIfMissing)
}

func Gatewaypluginaia2aproxyStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_a2a_proxy", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaiawsguardrailsStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_aws_guardrails", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaiazurecontentsafetyStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_azure_content_safety", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaicustomguardrailStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_custom_guardrail", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaigcpmodelarmorStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_gcp_model_armor", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginailakeraguardStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_lakera_guard", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaillmasjudgeStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_llm_as_judge", req, resp, defaultWorkspaceIfMissing)
}

func Gatewaypluginaimcpoauth2StateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_mcp_oauth2", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaimcpproxyStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_mcp_proxy", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaipromptcompressorStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_prompt_compressor", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaipromptdecoratorStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_prompt_decorator", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaipromptguardStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_prompt_guard", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaiprompttemplateStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_prompt_template", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaiproxyStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_proxy", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaiproxyadvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_proxy_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginairaginjectorStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_rag_injector", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginairatelimitingadvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_rate_limiting_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginairequesttransformerStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_request_transformer", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginairesponsetransformerStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_response_transformer", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaisanitizerStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_sanitizer", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaisemanticcacheStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_semantic_cache", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaisemanticpromptguardStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_semantic_prompt_guard", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginaisemanticresponseguardStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ai_semantic_response_guard", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginappdynamicsStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_app_dynamics", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginawslambdaStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_aws_lambda", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginazurefunctionsStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_azure_functions", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginbasicauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_basic_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginbotdetectionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_bot_detection", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugincanaryStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_canary", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginconfluentStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_confluent", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginconfluentconsumeStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_confluent_consume", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugincorrelationidStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_correlation_id", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugincorsStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_cors", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugindatadogStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_datadog", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugindatakitStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_datakit", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugindegraphqlStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_degraphql", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginexittransformerStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_exit_transformer", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginfilelogStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_file_log", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginforwardproxyStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_forward_proxy", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugingraphqlproxycacheadvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_graphql_proxy_cache_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugingraphqlratelimitingadvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_graphql_rate_limiting_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugingrpcgatewayStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_grpc_gateway", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugingrpcwebStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_grpc_web", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginheadercertauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_header_cert_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginhmacauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_hmac_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginhttplogStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_http_log", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugininjectionprotectionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_injection_protection", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginiprestrictionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ip_restriction", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginjqStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_jq", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginjsonthreatprotectionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_json_threat_protection", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginjwedecryptStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_jwe_decrypt", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginjwtStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_jwt", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginjwtsignerStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_jwt_signer", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginkafkaconsumeStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_kafka_consume", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginkafkalogStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_kafka_log", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginkafkaupstreamStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_kafka_upstream", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginkeyauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_key_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginldapauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ldap_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginldapauthadvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_ldap_auth_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginlogglyStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_loggly", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginmeteringandbillingStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_metering_and_billing", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginmockingStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_mocking", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginmtlsauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_mtls_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginoasvalidationStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_oas_validation", req, resp, defaultWorkspaceIfMissing)
}

func Gatewaypluginoauth2introspectionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_oauth2_introspection", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginopaStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_opa", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginopenidconnectStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_openid_connect", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginopentelemetryStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_opentelemetry", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginpostfunctionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_post_function", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginprefunctionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_pre_function", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginprometheusStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_prometheus", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginproxycacheStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_proxy_cache", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginproxycacheadvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_proxy_cache_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginratelimitingStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_rate_limiting", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginratelimitingadvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_rate_limiting_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginredirectStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_redirect", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginrequestcalloutStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_request_callout", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginrequestsizelimitingStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_request_size_limiting", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginrequestterminationStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_request_termination", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginrequesttransformerStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_request_transformer", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginrequesttransformeradvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_request_transformer_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginrequestvalidatorStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_request_validator", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginresponseratelimitingStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_response_ratelimiting", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginresponsetransformerStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_response_transformer", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginresponsetransformeradvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_response_transformer_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginroutebyheaderStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_route_by_header", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginroutetransformeradvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_route_transformer_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginsamlStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_saml", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginserviceprotectionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_service_protection", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginsessionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_session", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginsolaceconsumeStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_solace_consume", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginsolacelogStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_solace_log", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginsolaceupstreamStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_solace_upstream", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginstandardwebhooksStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_standard_webhooks", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginstatsdStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_statsd", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginstatsdadvancedStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_statsd_advanced", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginsyslogStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_syslog", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugintcplogStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_tcp_log", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugintlshandshakemodifierStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_tls_handshake_modifier", req, resp, defaultWorkspaceIfMissing)
}

func GatewayplugintlsmetadataheadersStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_tls_metadata_headers", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginudplogStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_udp_log", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginupstreamoauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_upstream_oauth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginupstreamtimeoutStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_upstream_timeout", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginvaultauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_vault_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginwebsocketsizelimitStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_websocket_size_limit", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginwebsocketvalidatorStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_websocket_validator", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginxmlthreatprotectionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_xml_threat_protection", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypluginzipkinStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_plugin_zipkin", req, resp, defaultWorkspaceIfMissing)
}

func GatewayrouteStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_route", req, resp, defaultWorkspaceIfMissing)
}

func GatewayrouteexpressionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_route_expression", req, resp, defaultWorkspaceIfMissing)
}

func GatewayserviceStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_service", req, resp, defaultWorkspaceIfMissing)
}

func GatewaysniStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_sni", req, resp, defaultWorkspaceIfMissing)
}

func GatewaytargetStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_target", req, resp, defaultWorkspaceIfMissing)
}

func GatewayupstreamStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_upstream", req, resp, defaultWorkspaceIfMissing)
}

func GatewayvaultStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_vault", req, resp, defaultWorkspaceIfMissing)
}
