package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func aiGatewayCase(check resource.TestCheckFunc) resource.TestCase {
	return resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV6ProviderFactories: providerFactory,
				ConfigDirectory:          config.TestNameDirectory(),
				Check:                    check,
			},
		},
	}
}

func TestAIGateway(t *testing.T) {
	t.Run("ai-gateway", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway.my_aigateway", "name", "tf-test-ai-gateway"),
		))
	})

	t.Run("data-plane-certificate", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_data_plane_certificate.my_aigatewaydataplanecertificate", "title", "tf-test-dp-cert"),
		))
	})

	t.Run("consumer", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_consumer.my_aigatewayconsumer", "name", "tf-test-consumer"),
		))
	})

	t.Run("consumer-group", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_consumer_group.my_aigatewayconsumergroup", "name", "tf-test-consumers"),
		))
	})

	t.Run("consumer-group-member", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttrSet("konnect_ai_gateway_consumer_group_member.my_aigatewayconsumergroupmember", "consumer_id"),
		))
	})

	t.Run("consumer-credential", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_consumer_credential.my_aigatewayconsumercredential", "display_name", "TF Test Dev Key"),
		))
	})

	t.Run("config-store", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_config_store.my_aigatewayconfigstore", "name", "tf-test-config-store"),
		))
	})

	t.Run("vault", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_vault.my_aigatewayvault", "name", "tf-test-vault"),
		))
	})

	t.Run("config-store-secret", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("konnect_ai_gateway_config_store_secret.my_aigatewayconfigstoresecret", "key", "tf-test-secret-key"),
				resource.TestCheckResourceAttr("konnect_ai_gateway_config_store_secret.my_aigatewayconfigstoresecret", "value", "tf-test-secret-value"),
			),
		))
	})

	t.Run("policy", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_policy.my_aigatewaypolicy", "name", "ai-pii-sanitizer-1234"),
		))
	})

	t.Run("model-provider-azure", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_model_provider.my_aigatewaymodelprovider", "name", "tf-test-azure-ai-provider"),
		))
	})

	t.Run("model-provider-anthropic", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_model_provider.my_aigatewaymodelprovider_anthropic", "name", "tf-test-anthropic-provider"),
		))
	})

	t.Run("model-anthropic", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_model.my_aigatewaymodel", "name", "tf-test-claude-5-model"),
		))
	})

	t.Run("model-azure", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_model.my_aigatewaymodel_model", "name", "my-azure-model"),
		))
	})

	t.Run("mcp-listener-upstream", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttrSet("konnect_ai_gateway_mcp_server.my_aigatewaymcpserver_listener", "id"),
		))
	})

	t.Run("mcp-conversion", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttrSet("konnect_ai_gateway_mcp_server.my_aigatewaymcpserver_conversion", "id"),
		))
	})

	t.Run("agent", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_agent.my_aigatewayagent", "name", "tf-test-flight-booking-agent"),
		))
	})

	t.Run("identity-providers", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("konnect_ai_gateway_identity_provider.my_aigatewayidentityprovider", "name", "tf-test-key-auth-identity-provider"),
				resource.TestCheckResourceAttr("konnect_ai_gateway_identity_provider.my_aigatewayidentityprovider_oidc", "name", "tf-test-openid-connect-identity-provider"),
			),
		))
	})

	t.Run("certificate", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_certificate.my_aigatewaycertificate", "name", "tf-test-certificate"),
		))
	})

	t.Run("ca-certificate", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_ca_certificate.my_aigatewaycacertificate", "name", "tf-test-ca-certificate"),
		))
	})

	t.Run("auth-strategy", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttrSet("konnect_ai_gateway_auth_strategy.my_aigatewayauthstrategy", "id"),
		))
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttrSet("konnect_ai_gateway_auth_strategy.my_aigatewayauthstrategy_oidc", "id"),
		))
	})

	t.Run("sni", func(t *testing.T) {
		resource.Test(t, aiGatewayCase(
			resource.TestCheckResourceAttr("konnect_ai_gateway_sni.my_aigatewaysni", "name", "tf-test-sni"),
		))
	})
}
