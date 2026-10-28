package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

// eventGatewayCPV13 pins the gateway to min_runtime_version 1.3, which the
// policy resources and fields exercised in this file require. The backend
// cluster config is shared from event_gateway_test.go (same package).
const eventGatewayCPV13 = `
	resource "konnect_event_gateway" "my_event_gateway" {
	  name                = "my_test_event_gateway"
	  min_runtime_version = "1.3"
	}
`

const eventGatewayVirtualClusterV13 = `
	resource "konnect_event_gateway_virtual_cluster" "my_event_gateway_virtual_cluster" {
	  name      = "my_test_event_gateway_virtual_cluster"
	  dns_label = "vcluster-1"
	  acl_mode  = "passthrough"
	  authentication = [
			{
			  anonymous = {}
			}
	  ]
	  destination = {
			id = konnect_event_gateway_backend_cluster.my_event_gateway_backend_cluster.id
	  }
	}
`

// egwV13Base is the scaffolding every subtest here shares: a gateway pinned to
// 1.3 plus a backend cluster and virtual cluster wired to it. Subtests add only
// the resources under test.
type egwV13Base struct {
	providerURL    string
	gateway        *hclbuilder.Builder
	backendCluster *hclbuilder.Builder
	virtualCluster *hclbuilder.Builder
}

func newEGWV13Base(t *testing.T, providerURL string) egwV13Base {
	t.Helper()

	gateway, err := hclbuilder.FromString(eventGatewayCPV13)
	require.NoError(t, err)

	backendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
	require.NoError(t, err)
	backendCluster.AddAttribute("gateway_id", gateway.ResourcePath()+".id")

	virtualCluster, err := hclbuilder.FromString(eventGatewayVirtualClusterV13)
	require.NoError(t, err)
	virtualCluster.AddAttribute("gateway_id", gateway.ResourcePath()+".id")

	return egwV13Base{
		providerURL:    providerURL,
		gateway:        gateway,
		backendCluster: backendCluster,
		virtualCluster: virtualCluster,
	}
}

// config renders the provider block, the shared scaffolding and the extra
// resources for a step.
func (b egwV13Base) config(extras ...*hclbuilder.Builder) string {
	builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, b.providerURL)
	builder.ProviderProperty = hclbuilder.Konnect
	out := builder.Upsert(b.gateway).Upsert(b.backendCluster).Upsert(b.virtualCluster)
	for _, extra := range extras {
		out = out.Upsert(extra)
	}
	return out.Build()
}

func TestEventGatewayV13Reusable(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerURL := fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort)

	t.Run("EGW Cluster Policy Request Rule Validator", func(t *testing.T) {
		const addr = "konnect_event_gateway_cluster_policy_request_rule_validator.test_request_rule_validator"

		base := newEGWV13Base(t, providerURL)

		requestRuleValidator, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_cluster_policy_request_rule_validator" "test_request_rule_validator" {
				name        = "test-request-rule-validator"
				description = "Test cluster policy for request rule validation"
				condition   = "context.auth.principal.name == 'team-a-service'"
				enabled     = true

				config = {
					requests = [
						{
							create_topics = {
								rules = [
									{
										action      = "reject"
										description = "Enforce the topic naming convention and a partition limit."
										rule        = "context.topic.name.startsWith('team-a.') && context.topic.num_partitions <= 50"
									}
								]
							}
						},
						{
							consumer_group_heartbeat = {
								rules = [
									{
										action      = "passthrough"
										description = "Log heartbeats sent by consumer groups outside the team-a prefix."
										rule        = "context.group.id.startsWith('team-a.')"
									}
								]
							}
						}
					]
				}
			}
		`)
		require.NoError(t, err)
		requestRuleValidator.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		requestRuleValidator.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: base.config(requestRuleValidator),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addr, "config.requests.0.create_topics.rules.0.action", "reject"),
						resource.TestCheckResourceAttr(addr, "config.requests.1.consumer_group_heartbeat.rules.0.action", "passthrough"),
					),
				},
			},
		})
	})

	t.Run("EGW Consume Policy Mask Fields", func(t *testing.T) {
		const addr = "konnect_event_gateway_consume_policy_mask_fields.test_mask_fields_policy"

		base := newEGWV13Base(t, providerURL)

		// mask_fields is a child of a schema_validation policy, so a parent
		// schema_validation policy is created and referenced via parent_policy_id.
		parentPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_schema_validation" "test_mask_fields_parent" {
				name        = "test-mask-fields-parent"
				description = "Parent schema validation policy for mask fields"
				condition   = "context.topic.name == 'validated-topic'"
				enabled     = true

				config = {
					json = {
						key_validation_action   = "mark"
						value_validation_action = "skip"
					}
				}
			}
		`)
		require.NoError(t, err)
		parentPolicy.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		parentPolicy.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")

		maskFieldsPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_mask_fields" "test_mask_fields_policy" {
				name        = "test-consume-mask-fields-policy"
				description = "Test consume policy for masking parsed record fields"
				condition   = "record.value.content.foo.bar == 'a-value'"
				enabled     = true

				config = {
					failure_mode = "mark"
					mask_fields = [
						{
							paths = [
								{
									match = "user.email"
								}
							]
							strategy = {
								email = {
									domain = {
										keep_chars = {
											first  = 1
											last   = 1
											phrase = "***"
										}
									}
									local_part = {
										replace = {
											phrase = "[redacted]"
										}
									}
								}
							}
						},
						{
							paths = [
								{
									match = "user.phone"
								}
							]
							strategy = {
								keep_chars = {
									first  = 0
									last   = 4
									phrase = "***"
								}
							}
						}
					]
				}
			}
		`)
		require.NoError(t, err)
		maskFieldsPolicy.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		maskFieldsPolicy.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")
		maskFieldsPolicy.AddAttribute("parent_policy_id", parentPolicy.ResourcePath()+".id")

		baseConfig := base.config(parentPolicy, maskFieldsPolicy)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: baseConfig,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addr, "config.mask_fields.0.paths.0.match", "user.email"),
						resource.TestCheckResourceAttr(addr, "config.mask_fields.0.strategy.email.domain.keep_chars.phrase", "***"),
						resource.TestCheckResourceAttr(addr, "config.mask_fields.1.strategy.keep_chars.last", "4"),
					),
				},
				{
					Config: base.config(parentPolicy, maskFieldsPolicy.AddAttribute("labels", `{ env = "test", policy = "mask-fields" }`)),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(addr, plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addr, "labels.env", "test"),
					),
				},
			},
		})
	})

	t.Run("EGW Produce Policy Mask Fields", func(t *testing.T) {
		const addr = "konnect_event_gateway_produce_policy_mask_fields.test_mask_fields_policy"

		base := newEGWV13Base(t, providerURL)

		parentPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_schema_validation" "test_mask_fields_parent" {
				name        = "test-mask-fields-parent"
				description = "Parent schema validation policy for mask fields"
				enabled     = true

				config = {
					json = {
						key_validation_action   = "mark"
						value_validation_action = "reject"
					}
				}
			}
		`)
		require.NoError(t, err)
		parentPolicy.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		parentPolicy.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")

		maskFieldsPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_mask_fields" "test_mask_fields_policy" {
				name        = "test-produce-mask-fields-policy"
				description = "Test produce policy for masking parsed record fields"
				condition   = "record.value.content.foo.bar == 'a-value'"
				enabled     = true

				config = {
					failure_mode = "passthrough"
					mask_fields = [
						{
							paths = [
								{
									match = "user.card.number"
								}
							]
							strategy = {
								replace = {
									phrase = "[redacted]"
								}
							}
						}
					]
				}
			}
		`)
		require.NoError(t, err)
		maskFieldsPolicy.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		maskFieldsPolicy.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")
		maskFieldsPolicy.AddAttribute("parent_policy_id", parentPolicy.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: base.config(parentPolicy, maskFieldsPolicy),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addr, "config.mask_fields.0.paths.0.match", "user.card.number"),
						resource.TestCheckResourceAttr(addr, "config.mask_fields.0.strategy.replace.phrase", "[redacted]"),
					),
				},
			},
		})
	})

	t.Run("EGW Consume Policy Transcode", func(t *testing.T) {
		const addr = "konnect_event_gateway_consume_policy_transcode.test_transcode_policy"

		base := newEGWV13Base(t, providerURL)

		// transcode is a child of a schema_validation policy, so a parent
		// schema_validation policy is created and referenced via parent_policy_id.
		parentPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_schema_validation" "test_transcode_parent" {
				name        = "test-transcode-parent"
				description = "Parent schema validation policy for transcode"
				condition   = "context.topic.name == 'validated-topic'"
				enabled     = true

				config = {
					json = {
						key_validation_action   = "mark"
						value_validation_action = "skip"
					}
				}
			}
		`)
		require.NoError(t, err)
		parentPolicy.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		parentPolicy.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")

		transcodePolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_transcode" "test_transcode_policy" {
				name        = "test-consume-transcode-policy"
				description = "Test consume policy for transcoding parsed records to JSON"
				condition   = "record.value.content.foo.bar == 'a-value'"
				enabled     = true

				config = {
					json = {
						failure_mode     = "mark"
						on_unknown_field = "drop"
						schema_source = {
							inline = {
								inline = "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"}}}"
							}
						}
					}
				}
			}
		`)
		require.NoError(t, err)
		transcodePolicy.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		transcodePolicy.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")
		transcodePolicy.AddAttribute("parent_policy_id", parentPolicy.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: base.config(parentPolicy, transcodePolicy),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addr, "config.json.on_unknown_field", "drop"),
						resource.TestCheckResourceAttr(addr, "config.json.schema_source.inline.inline", `{"type":"object","properties":{"id":{"type":"string"}}}`),
					),
				},
			},
		})
	})

	t.Run("EGW Produce Policy Transcode", func(t *testing.T) {
		const addr = "konnect_event_gateway_produce_policy_transcode.test_transcode_policy"

		base := newEGWV13Base(t, providerURL)

		// The Avro transcode policy requires a schema_source, which here refers
		// to a schema registered in a schema registry.
		schemaRegistry, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_schema_registry" "test_transcode_registry" {
				confluent = {
					name        = "transcode-schema-registry"
					description = "schema registry for produce transcode policy"

					config = {
						endpoint        = "https://key-hovercraft.com"
						schema_type     = "avro"
						timeout_seconds = 8

						authentication = {
							basic = {
								username = "test-user"
								password = "$${vault.env['MY_ENV_VAR']}"
							}
						}
					}
				}
			}
		`)
		require.NoError(t, err)
		schemaRegistry.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")

		parentPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_schema_validation" "test_transcode_parent" {
				name        = "test-transcode-parent"
				description = "Parent schema validation policy for transcode"
				enabled     = true

				config = {
					confluent_schema_registry = {
						key_validation_action   = "reject"
						value_validation_action = "reject"
						schema_registry = {
							id = konnect_event_gateway_schema_registry.test_transcode_registry.id
						}
					}
				}
			}
		`)
		require.NoError(t, err)
		parentPolicy.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		parentPolicy.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")
		parentPolicy.AddAttribute(
			"depends_on",
			`[konnect_event_gateway_schema_registry.test_transcode_registry]`,
		)

		transcodePolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_transcode" "test_transcode_policy" {
				name        = "test-produce-transcode-policy"
				description = "Test produce policy for transcoding parsed records to Avro"
				condition   = "record.value.content.foo.bar == 'a-value'"
				enabled     = true

				config = {
					avro = {
						failure_mode     = "mark"
						on_unknown_field = "drop"
						schema_source = {
							reference = {
								schema_ref_destination = {
									confluent_format = {}
								}
								schema_registry = {
									id = konnect_event_gateway_schema_registry.test_transcode_registry.id
								}
								subject = "context.topic.name + '-value'"
								version = "'latest'"
							}
						}
					}
				}
			}
		`)
		require.NoError(t, err)
		transcodePolicy.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		transcodePolicy.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")
		transcodePolicy.AddAttribute("parent_policy_id", parentPolicy.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: base.config(schemaRegistry, parentPolicy, transcodePolicy),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addr, "config.avro.on_unknown_field", "drop"),
						resource.TestCheckResourceAttr(addr, "config.avro.schema_source.reference.subject", "context.topic.name + '-value'"),
						resource.TestCheckResourceAttr(addr, "config.avro.schema_source.reference.version", "'latest'"),
					),
				},
			},
		})
	})

	t.Run("EGW Consume Policy Schema Validation Inline Schema", func(t *testing.T) {
		const addr = "konnect_event_gateway_consume_policy_schema_validation.test_inline_schema_policy"

		base := newEGWV13Base(t, providerURL)

		inlineSchemaPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_schema_validation" "test_inline_schema_policy" {
				name        = "test-consume-inline-schema-policy"
				description = "Test consume policy for inline schema validation"
				condition   = "context.topic.name == 'validated-topic'"
				enabled     = true

				config = {
					inline_schema = {
						failure_mode   = "mark"
						validate_key   = false
						validate_value = true

						value = {
							json = {
								schema = "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"}}}"
							}
						}
					}
				}
			}
		`)
		require.NoError(t, err)
		inlineSchemaPolicy.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		inlineSchemaPolicy.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: base.config(inlineSchemaPolicy),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addr, "config.inline_schema.value.json.schema", `{"type":"object","properties":{"id":{"type":"string"}}}`),
					),
				},
			},
		})
	})

	t.Run("EGW Produce Policy Schema Validation Inline Schema", func(t *testing.T) {
		const addr = "konnect_event_gateway_produce_policy_schema_validation.test_inline_schema_policy"

		base := newEGWV13Base(t, providerURL)

		inlineSchemaPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_schema_validation" "test_inline_schema_policy" {
				name        = "test-produce-inline-schema-policy"
				description = "Test produce policy for inline schema validation"
				enabled     = true

				config = {
					inline_schema = {
						failure_mode   = "reject"
						validate_key   = false
						validate_value = true

						value = {
							avro = {
								schema = "{\"type\":\"record\",\"name\":\"Order\",\"fields\":[{\"name\":\"id\",\"type\":\"string\"}]}"
							}
						}
					}
				}
			}
		`)
		require.NoError(t, err)
		inlineSchemaPolicy.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")
		inlineSchemaPolicy.AddAttribute("virtual_cluster_id", base.virtualCluster.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: base.config(inlineSchemaPolicy),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addr, "config.inline_schema.value.avro.schema", `{"type":"record","name":"Order","fields":[{"name":"id","type":"string"}]}`),
					),
				},
			},
		})
	})

	t.Run("EGW Backend Cluster SASL AWS IAM", func(t *testing.T) {
		const addr = "konnect_event_gateway_backend_cluster.test_sasl_aws_iam_backend_cluster"

		base := newEGWV13Base(t, providerURL)

		backendCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_backend_cluster" "test_sasl_aws_iam_backend_cluster" {
				name = "test-sasl-aws-iam-backend-cluster"
				authentication = {
					sasl_aws_iam = {
						sasl_aws_iam = {
							assume_role = {
								arn          = "arn:aws:iam::123456789012:role/msk-consumer"
								session_name = "kong-egw"
							}
						}
					}
				}
				bootstrap_servers = [
					"b-1.example-cluster.kafka.us-east-1.amazonaws.com:9098"
				]
				tls = {
					enabled              = true
					insecure_skip_verify = false
				}
			}
		`)
		require.NoError(t, err)
		backendCluster.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: base.config(backendCluster),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addr, "authentication.sasl_aws_iam.sasl_aws_iam.assume_role.arn", "arn:aws:iam::123456789012:role/msk-consumer"),
						resource.TestCheckResourceAttr(addr, "authentication.sasl_aws_iam.sasl_aws_iam.assume_role.session_name", "kong-egw"),
					),
				},
			},
		})
	})

	t.Run("EGW Schema Registry Any Schema Type", func(t *testing.T) {
		const addr = "konnect_event_gateway_schema_registry.test_any_schema_registry"

		base := newEGWV13Base(t, providerURL)

		schemaRegistry, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_schema_registry" "test_any_schema_registry" {
				confluent = {
					name        = "test-any-schema-registry"
					description = "test schema registry accepting any schema type"

					config = {
						endpoint        = "https://key-hovercraft.com"
						schema_type     = "any"
						timeout_seconds = 8

						authentication = {
							basic = {
								username = "test-user"
								password = "$${vault.env['MY_ENV_VAR']}"
							}
						}
					}
				}
			}
		`)
		require.NoError(t, err)
		schemaRegistry.AddAttribute("gateway_id", base.gateway.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: base.config(schemaRegistry),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addr, "confluent.config.schema_type", "any"),
					),
				},
			},
		})
	})
}
