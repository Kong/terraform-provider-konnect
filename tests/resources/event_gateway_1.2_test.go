package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

// eventGatewayCPV12 pins the gateway to min_runtime_version 1.2, which the
// policy resources exercised in this file require. The backend-cluster config
// is shared from event_gateway_test.go (same package).
const eventGatewayCPV12 = `
	resource "konnect_event_gateway" "my_event_gateway" {
	  name                = "my_test_event_gateway"
	  min_runtime_version = "1.2"
	}
`

const eventGatewayVirtualClusterV12 = `
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
	  topic_aliases = [
			{
			  alias    = "my-alias-topic"
			  topic    = "my-backend-topic"
			  conflict = "warn"
			}
	  ]
	}
`

func TestEventGatewayV12Reusable(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("EGW Produce Policy Encrypt", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCPV12)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualClusterV12)
		require.NoError(t, err)

		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		staticKey, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_static_key" "test_encryption_key" {
				name  = "test-encryption-key"
				value = "$${vault.env[\"MY_ENV_VAR\"]}"
			}
		`)
		require.NoError(t, err)

		staticKey.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		producePolicyEncrypt, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_encrypt" "test_produce_encrypt_policy" {
				name        = "test-produce-encrypt-policy"
				description = "Test produce policy for encryption"
				condition   = "context.topic.name == 'encrypted-topic'"
				enabled     = true

				config = {
					failure_mode = "error"
					encryption_key = {
					  static = {
						key = {
						  id = konnect_event_gateway_static_key.test_encryption_key.id
						}
					  }
					}
					part_of_record = ["value"]
				}
			}
		`)
		require.NoError(t, err)

		producePolicyEncrypt.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)
		producePolicyEncrypt.AddAttribute(
			"virtual_cluster_id",
			egwVirtualCluster.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(staticKey).Upsert(producePolicyEncrypt).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", "name", "test-produce-encrypt-policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", "description", "Test produce policy for encryption"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", "enabled", "true")),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(staticKey).Upsert(producePolicyEncrypt).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(staticKey).Upsert(producePolicyEncrypt.AddAttribute("description", "Updated Test produce policy for encryption")).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", "description", "Updated Test produce policy for encryption")),
				},
			},
		})
	})

	t.Run("EGW Produce Policy Encrypt Fields", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCPV12)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)
		egwBackendCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualClusterV12)
		require.NoError(t, err)
		egwVirtualCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		schemaRegistry, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_schema_registry" "test_encrypt_fields_registry" {
			  confluent = {
				name        = "encrypt-fields-schema-registry"
				description = "schema registry for encrypt fields parent policy"

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
		schemaRegistry.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		staticKey, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_static_key" "test_encrypt_fields_key" {
				name  = "test-encrypt-fields-key"
				value = "$${vault.env[\"MY_ENV_VAR\"]}"
			}
		`)
		require.NoError(t, err)
		staticKey.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		// encrypt_fields is a child of a schema_validation policy, so a parent
		// schema_validation policy is created and referenced via parent_policy_id.
		parentPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_schema_validation" "test_encrypt_fields_parent" {
				name        = "test-encrypt-fields-parent"
				description = "Parent schema validation policy for encrypt fields"
				enabled     = true

				config = {
					confluent_schema_registry = {
						key_validation_action   = "reject"
						value_validation_action = "reject"
						schema_registry = {
							id = konnect_event_gateway_schema_registry.test_encrypt_fields_registry.id
						}
					}
				}
			}
		`)
		require.NoError(t, err)
		parentPolicy.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		parentPolicy.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")
		parentPolicy.AddAttribute(
			"depends_on",
			`[konnect_event_gateway_schema_registry.test_encrypt_fields_registry]`,
		)

		encryptFieldsPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_encrypt_fields" "test_encrypt_fields_policy" {
				name        = "test-produce-encrypt-fields-policy"
				description = "Test produce policy for encrypting parsed record fields"
				condition   = "record.value.content.foo.bar == 'a-value'"
				enabled     = true

				config = {
					failure_mode = "mark"
					encrypt_fields = [
						{
							encryption_key = {
								static = {
									key = {
										id = konnect_event_gateway_static_key.test_encrypt_fields_key.id
									}
								}
							}
							paths = [
								{
									match = "someObject.someArray[1].fieldName"
								}
							]
						}
					]
				}
			}
		`)
		require.NoError(t, err)
		encryptFieldsPolicy.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		encryptFieldsPolicy.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")
		encryptFieldsPolicy.AddAttribute("parent_policy_id", parentPolicy.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(schemaRegistry).Upsert(staticKey).Upsert(parentPolicy).Upsert(encryptFieldsPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_produce_policy_encrypt_fields.test_encrypt_fields_policy", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt_fields.test_encrypt_fields_policy", "name", "test-produce-encrypt-fields-policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt_fields.test_encrypt_fields_policy", "description", "Test produce policy for encrypting parsed record fields"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt_fields.test_encrypt_fields_policy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt_fields.test_encrypt_fields_policy", "config.failure_mode", "mark"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt_fields.test_encrypt_fields_policy", "config.encrypt_fields.0.paths.0.match", "someObject.someArray[1].fieldName"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(schemaRegistry).Upsert(staticKey).Upsert(parentPolicy).Upsert(encryptFieldsPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(schemaRegistry).Upsert(staticKey).Upsert(parentPolicy).Upsert(encryptFieldsPolicy.AddAttribute("description", "Updated Test produce policy for encrypting parsed record fields")).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt_fields.test_encrypt_fields_policy", "description", "Updated Test produce policy for encrypting parsed record fields"),
					),
				},
			},
		})
	})

	t.Run("EGW Consume Policy Decrypt", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCPV12)
		require.NoError(t, err)

		backend, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		backend.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		virtualCluster, err := hclbuilder.FromString(eventGatewayVirtualClusterV12)
		require.NoError(t, err)

		virtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		staticKey, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_static_key" "tf-test-static-key" {
			  name  = "tf-test-static-key"
			  value = "$${vault.env[\"MY_ENV_VAR\"]}"
			}
		`)
		require.NoError(t, err)

		staticKey.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		decryptPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_decrypt" "tf-test-decrypt-policy" {
			  name        = "tf-test-decrypt-policy"
			  description = "initial decrypt policy"
			  enabled     = false
			  condition   = "context.topic.name.endsWith('my_suffix')"

			  config = {
				failure_mode = "passthrough"

				key_sources = [
				  {
					static = {
					  name = "tf-test-static-key"
					}
				  }
				]
				part_of_record = ["key"]
			  }
			}
		`)
		require.NoError(t, err)

		decryptPolicy.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)
		decryptPolicy.AddAttribute(
			"virtual_cluster_id",
			virtualCluster.ResourcePath()+".id",
		)
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(backend).Upsert(virtualCluster).Upsert(staticKey).Upsert(decryptPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_consume_policy_decrypt.tf-test-decrypt-policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt.tf-test-decrypt-policy", "enabled", "false"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt.tf-test-decrypt-policy", "config.failure_mode", "passthrough"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(backend).Upsert(virtualCluster).Upsert(staticKey).Upsert(decryptPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(backend).Upsert(virtualCluster).Upsert(staticKey).Upsert(decryptPolicy.AddAttribute("enabled", "true")).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_consume_policy_decrypt.tf-test-decrypt-policy", plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_decrypt.tf-test-decrypt-policy",
							"enabled",
							"true",
						)),
				},
			},
		})
	})

	t.Run("EGW Consume Policy Decrypt Fields", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCPV12)
		require.NoError(t, err)

		backend, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)
		backend.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		virtualCluster, err := hclbuilder.FromString(eventGatewayVirtualClusterV12)
		require.NoError(t, err)
		virtualCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		// The decrypt_fields policy uses key_sources with an empty `static = {}`,
		// which resolves against static keys present in the gateway, so at least
		// one must exist or the API rejects the policy with a 400.
		staticKey, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_static_key" "test_decrypt_fields_key" {
				name  = "test-decrypt-fields-key"
				value = "$${vault.env[\"MY_ENV_VAR\"]}"
			}
		`)
		require.NoError(t, err)
		staticKey.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		// decrypt_fields is a child of a schema_validation policy, so a parent
		// schema_validation policy is created and referenced via parent_policy_id.
		parentPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_schema_validation" "test_decrypt_fields_parent" {
				name        = "test-decrypt-fields-parent"
				description = "Parent schema validation policy for decrypt fields"
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
		parentPolicy.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		parentPolicy.AddAttribute("virtual_cluster_id", virtualCluster.ResourcePath()+".id")

		decryptFieldsPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_decrypt_fields" "test_decrypt_fields_policy" {
				name        = "test-consume-decrypt-fields-policy"
				description = "Test consume policy for decrypting parsed record fields"
				condition   = "record.value.content.foo.bar == 'a-value'"
				enabled     = true

				config = {
					failure_mode = "mark"
					key_sources = [
						{
							static = {}
						}
					]
					decrypt_fields = {
						paths = [
							{
								match = "someObject.someArray[1].fieldName"
							}
						]
					}
				}
			}
		`)
		require.NoError(t, err)
		decryptFieldsPolicy.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		decryptFieldsPolicy.AddAttribute("virtual_cluster_id", virtualCluster.ResourcePath()+".id")
		decryptFieldsPolicy.AddAttribute("parent_policy_id", parentPolicy.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(backend).Upsert(virtualCluster).Upsert(staticKey).Upsert(parentPolicy).Upsert(decryptFieldsPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_consume_policy_decrypt_fields.test_decrypt_fields_policy", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt_fields.test_decrypt_fields_policy", "name", "test-consume-decrypt-fields-policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt_fields.test_decrypt_fields_policy", "description", "Test consume policy for decrypting parsed record fields"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt_fields.test_decrypt_fields_policy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt_fields.test_decrypt_fields_policy", "config.failure_mode", "mark"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt_fields.test_decrypt_fields_policy", "config.decrypt_fields.paths.0.match", "someObject.someArray[1].fieldName"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(backend).Upsert(virtualCluster).Upsert(staticKey).Upsert(parentPolicy).Upsert(decryptFieldsPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(backend).Upsert(virtualCluster).Upsert(staticKey).Upsert(parentPolicy).Upsert(decryptFieldsPolicy.AddAttribute("labels", `{ env = "test", policy = "decrypt-fields" }`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt_fields.test_decrypt_fields_policy", "labels.%", "2"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt_fields.test_decrypt_fields_policy", "labels.env", "test"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt_fields.test_decrypt_fields_policy", "labels.policy", "decrypt-fields"),
					),
				},
			},
		})
	})
}
