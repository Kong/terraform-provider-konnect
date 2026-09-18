package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAIGateway(t *testing.T) {
	t.Run("ai-gateway", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway.my_aigateway", "name", "tf-test-ai-gateway"),
						resource.TestCheckResourceAttr("konnect_ai_gateway.my_aigateway", "display_name", "TF Test AI Gateway"),
					),
				},
				{
					// Update display_name
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway.my_aigateway", "display_name", "TF Test AI Gateway Updated"),
					),
				},
			},
		})
	})

	t.Run("data-plane-certificate", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_data_plane_certificate.my_aigatewaydataplanecertificate", "title", "tf-test-dp-cert"),
					),
				},
				{
					// Update title
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_data_plane_certificate.my_aigatewaydataplanecertificate", "title", "tf-test-dp-cert-updated"),
					),
				},
			},
		})
	})

	t.Run("consumer", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer.my_aigatewayconsumer", "name", "tf-test-consumer"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer.my_aigatewayconsumer", "display_name", "TF Test Consumer"),
					),
				},
				{
					// Update display_name
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer.my_aigatewayconsumer", "display_name", "TF Test Consumer Updated"),
					),
				},
			},
		})
	})

	t.Run("consumer-group", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer_group.my_aigatewayconsumergroup", "name", "tf-test-consumers"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer_group.my_aigatewayconsumergroup", "display_name", "TF Test Consumer Group"),
					),
				},
				{
					// Update display_name
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer_group.my_aigatewayconsumergroup", "display_name", "TF Test Consumer Group Updated"),
					),
				},
			},
		})
	})

	t.Run("consumer-group-member", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_ai_gateway_consumer_group_member.my_aigatewayconsumergroupmember", "consumer_id"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer.my_aigatewayconsumer", "display_name", "TF Test Consumer"),
					),
				},
				{
					// The member resource itself only references immutable IDs, so the
					// update step exercises an in-place change on the sibling consumer
					// and confirms the membership survives untouched.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer.my_aigatewayconsumer", "display_name", "TF Test Consumer Updated"),
						resource.TestCheckResourceAttrPair(
							"konnect_ai_gateway_consumer_group_member.my_aigatewayconsumergroupmember", "consumer_id",
							"konnect_ai_gateway_consumer.my_aigatewayconsumer", "id",
						),
					),
				},
			},
		})
	})

	t.Run("consumer-credential", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer_credential.my_aigatewayconsumercredential", "display_name", "TF Test Dev Key"),
					),
				},
				{
					// Update display_name
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer_credential.my_aigatewayconsumercredential", "display_name", "TF Test Dev Key Updated"),
					),
				},
			},
		})
	})

	t.Run("config-store", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_config_store.my_aigatewayconfigstore", "name", "tf-test-config-store"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_config_store.my_aigatewayconfigstore", "force", "true"),
					),
				},
				{
					// Update force
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_config_store.my_aigatewayconfigstore", "force", "false"),
					),
				},
			},
		})
	})

	t.Run("config-store-secret", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_config_store_secret.my_aigatewayconfigstoresecret", "key", "tf-test-secret-key"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_config_store_secret.my_aigatewayconfigstoresecret", "value", "tf-test-secret-value"),
					),
				},
				{
					// Update value
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_config_store_secret.my_aigatewayconfigstoresecret", "value", "tf-test-secret-value-updated"),
					),
				},
			},
		})
	})

	t.Run("vault", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_vault.my_aigatewayvault", "konnect.name", "tf-test-vault"),
					),
				},
				{
					// Update konnect.name
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_vault.my_aigatewayvault", "konnect.name", "tf-test-vault-updated"),
					),
				},
			},
		})
	})

	t.Run("policy", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_policy.my_aigatewaypolicy", "name", "ai-pii-sanitizer-1234"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_policy.my_aigatewaypolicy", "display_name", "My Cool AI PII Sanitizer Policy"),
					),
				},
				{
					// Update display_name
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_policy.my_aigatewaypolicy", "display_name", "My Cool AI PII Sanitizer Policy Updated"),
					),
				},
			},
		})
	})

	t.Run("model-provider", func(t *testing.T) {
		t.Run("anthropic", func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: providerFactory,
				Steps: []resource.TestStep{
					{
						Config:          providerConfigUs,
						ConfigDirectory: config.TestNameDirectory(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_model_provider.my_aigatewaymodelprovider_anthropic", "anthropic.name", "tf-test-anthropic-provider"),
							resource.TestCheckResourceAttr("konnect_ai_gateway_model_provider.my_aigatewaymodelprovider_anthropic", "anthropic.display_name", "TF Test Anthropic AI Provider"),
						),
					},
					{
						// Update display_name
						Config:          providerConfigUs,
						ConfigDirectory: config.TestStepDirectory(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_model_provider.my_aigatewaymodelprovider_anthropic", "anthropic.display_name", "TF Test Anthropic AI Provider Updated"),
						),
					},
				},
			})
		})

		t.Run("azure", func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: providerFactory,
				Steps: []resource.TestStep{
					{
						Config:          providerConfigUs,
						ConfigDirectory: config.TestNameDirectory(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_model_provider.my_aigatewaymodelprovider", "azure.name", "tf-test-azure-ai-provider"),
							resource.TestCheckResourceAttr("konnect_ai_gateway_model_provider.my_aigatewaymodelprovider", "azure.display_name", "Test TF Azure AI SE"),
						),
					},
					{
						// Update display_name
						Config:          providerConfigUs,
						ConfigDirectory: config.TestStepDirectory(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_model_provider.my_aigatewaymodelprovider", "azure.display_name", "Test TF Azure AI SE Updated"),
						),
					},
				},
			})
		})
	})

	t.Run("model", func(t *testing.T) {
		t.Run("anthropic", func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: providerFactory,
				Steps: []resource.TestStep{
					{
						Config:          providerConfigUs,
						ConfigDirectory: config.TestNameDirectory(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_model.my_aigatewaymodel", "api.name", "tf-test-claude-5-model"),
							resource.TestCheckResourceAttr("konnect_ai_gateway_model.my_aigatewaymodel", "api.display_name", "My Test Claude 5 model"),
						),
					},
					{
						// Update api.display_name
						Config:          providerConfigUs,
						ConfigDirectory: config.TestStepDirectory(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_model.my_aigatewaymodel", "api.display_name", "My Test Claude 5 model Updated"),
						),
					},
				},
			})
		})

		t.Run("azure", func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: providerFactory,
				Steps: []resource.TestStep{
					{
						Config:          providerConfigUs,
						ConfigDirectory: config.TestNameDirectory(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_model.my_aigatewaymodel_model", "model.name", "my-azure-model"),
							resource.TestCheckResourceAttr("konnect_ai_gateway_model.my_aigatewaymodel_model", "model.display_name", "My Test Azure model"),
						),
					},
					{
						// Update model.display_name
						Config:          providerConfigUs,
						ConfigDirectory: config.TestStepDirectory(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_model.my_aigatewaymodel_model", "model.display_name", "My Test Azure model Updated"),
						),
					},
				},
			})
		})
	})

	t.Run("mcp-server", func(t *testing.T) {
		t.Run("listener-upstream", func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: providerFactory,
				Steps: []resource.TestStep{
					{
						Config:          providerConfigUs,
						ConfigDirectory: config.TestNameDirectory(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttrSet("konnect_ai_gateway_mcp_server.my_aigatewaymcpserver_listener", "id"),
							resource.TestCheckResourceAttr("konnect_ai_gateway_mcp_server.my_aigatewaymcpserver_listener", "listener.display_name", "TF Test MCP Listener"),
						),
					},
					{
						// Update listener.display_name
						Config:          providerConfigUs,
						ConfigDirectory: config.TestStepDirectory(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_mcp_server.my_aigatewaymcpserver_listener", "listener.display_name", "TF Test MCP Listener Updated"),
						),
					},
				},
			})
		})

		t.Run("conversion", func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: providerFactory,
				Steps: []resource.TestStep{
					{
						Config:          providerConfigUs,
						ConfigDirectory: config.TestNameDirectory(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttrSet("konnect_ai_gateway_mcp_server.my_aigatewaymcpserver_conversion", "id"),
							resource.TestCheckResourceAttr("konnect_ai_gateway_mcp_server.my_aigatewaymcpserver_conversion", "conversion_only.display_name", "TF Test MCP Conversion"),
						),
					},
					{
						// Update conversion_only.display_name
						Config:          providerConfigUs,
						ConfigDirectory: config.TestStepDirectory(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_mcp_server.my_aigatewaymcpserver_conversion", "conversion_only.display_name", "TF Test MCP Conversion Updated"),
						),
					},
				},
			})
		})
	})

	t.Run("agent", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_agent.my_aigatewayagent", "name", "tf-test-flight-booking-agent"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_agent.my_aigatewayagent", "display_name", "Test TF Flight Booking Agent"),
					),
				},
				{
					// Update display_name
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_agent.my_aigatewayagent", "display_name", "Test TF Flight Booking Agent Updated"),
					),
				},
			},
		})
	})

	t.Run("certificate", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			Steps: []resource.TestStep{
				{
					ProtoV6ProviderFactories: providerFactory,
					Config:                   providerConfigUs,
					ConfigDirectory:          config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_certificate.my_aigatewaycertificate", "name", "tf-test-certificate"),
					),
				},
				{
					// Update name
					ProtoV6ProviderFactories: providerFactory,
					Config:                   providerConfigUs,
					ConfigDirectory:          config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_certificate.my_aigatewaycertificate", "labels.new", "my-label"),
					),
				},
			},
		})
	})

	t.Run("ca-certificate", func(t *testing.T) {
		resource.Test(t, resource.TestCase{

			Steps: []resource.TestStep{
				{
					ProtoV6ProviderFactories: providerFactory,
					ConfigDirectory:          config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_ca_certificate.my_aigatewaycacertificate", "name", "tf-test-ca-certificate"),
					),
				},
				{
					ProtoV6ProviderFactories: providerFactory,
					ConfigDirectory:          config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_ca_certificate.my_aigatewaycacertificate", "labels.new", "my-label"),
					),
				},
			},
		})
	})

	t.Run("auth-strategy", func(t *testing.T) {
		t.Run("key-auth", func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: providerFactory,
				Steps: []resource.TestStep{
					{
						Config:          providerConfigUs,
						ConfigDirectory: config.TestNameDirectory(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttrSet("konnect_ai_gateway_auth_strategy.my_aigatewayauthstrategy", "id"),
							resource.TestCheckResourceAttr("konnect_ai_gateway_auth_strategy.my_aigatewayauthstrategy", "key_auth.display_name", "TF Test Key Auth"),
						),
					},
					{
						// Update key_auth.display_name
						Config:          providerConfigUs,
						ConfigDirectory: config.TestStepDirectory(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_auth_strategy.my_aigatewayauthstrategy", "key_auth.display_name", "TF Test Key Auth Updated"),
						),
					},
				},
			})
		})

		t.Run("oidc", func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: providerFactory,
				Steps: []resource.TestStep{
					{
						Config:          providerConfigUs,
						ConfigDirectory: config.TestNameDirectory(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttrSet("konnect_ai_gateway_auth_strategy.my_aigatewayauthstrategy_oidc", "id"),
							resource.TestCheckResourceAttr("konnect_ai_gateway_auth_strategy.my_aigatewayauthstrategy_oidc", "openid_connect.display_name", "TF Test OpenID Connect"),
						),
					},
					{
						// Update openid_connect.display_name
						Config:          providerConfigUs,
						ConfigDirectory: config.TestStepDirectory(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_ai_gateway_auth_strategy.my_aigatewayauthstrategy_oidc", "openid_connect.display_name", "TF Test OpenID Connect Updated"),
						),
					},
				},
			})
		})
	})

	t.Run("sni", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			Steps: []resource.TestStep{
				{
					ProtoV6ProviderFactories: providerFactory,
					Config:                   providerConfigUs,
					ConfigDirectory:          config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_sni.my_aigatewaysni", "name", "tf-test-sni"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_sni.my_aigatewaysni", "display_name", "TF Test SNI"),
					),
				},
				{
					// Update display_name
					ProtoV6ProviderFactories: providerFactory,
					Config:                   providerConfigUs,
					ConfigDirectory:          config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway_sni.my_aigatewaysni", "display_name", "TF Test SNI Updated"),
					),
				},
			},
		})
	})
}
