package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

const (
	eventGatewayCP = `
		resource "konnect_event_gateway" "my_event_gateway" {
		  name = "my_test_event_gateway"
		}
	`
	eventGatewayBackendCluster = `
		resource "konnect_event_gateway_backend_cluster" "my_event_gateway_backend_cluster" {
		  name = "my_test_event_gateway_backend_cluster"
		  authentication = {
		    anonymous = {}
		  }
		  bootstrap_servers = [
		    "10.0.0.1:8080"
		  ]
		   tls = {
			  enabled = false
			  insecure_skip_verify = true
		   }
		}
	`
	eventGatewayVirtualCluster = `
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
)

func TestEventGatewayReusable(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("EGW Control plane", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.my_event_gateway", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway.my_event_gateway", "name", "my_test_event_gateway"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp.AddAttribute("labels", `{key = "value"}`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway.my_event_gateway", "labels.%", "1"),
						resource.TestCheckResourceAttr("konnect_event_gateway.my_event_gateway", "labels.key", "value"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})

	})

	t.Run("EGW Backend Cluster", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		egwBackendCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_backend_cluster.my_event_gateway_backend_cluster", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_backend_cluster.my_event_gateway_backend_cluster", "name", "my_test_event_gateway_backend_cluster"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster.AddAttribute("labels", `{key = "value"}`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_backend_cluster.my_event_gateway_backend_cluster", "labels.%", "1"),
						resource.TestCheckResourceAttr("konnect_event_gateway_backend_cluster.my_event_gateway_backend_cluster", "labels.key", "value"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("EGW Virtual Cluster", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)
		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)

		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_virtual_cluster.my_event_gateway_virtual_cluster",
								plancheck.ResourceActionCreate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster.my_event_gateway_virtual_cluster",
							"name",
							"my_test_event_gateway_virtual_cluster",
						),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(
						egwVirtualCluster.AddAttribute("labels", `{ key = "value" }`),
					).Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster.my_event_gateway_virtual_cluster",
							"labels.%",
							"1",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster.my_event_gateway_virtual_cluster",
							"labels.key",
							"value",
						),
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

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		backend, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		backend.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		virtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
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

	t.Run("EGW Consume Policy Modify Headers", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)

		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		consumePolicyModifyHeaders, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_modify_headers" "test_modify_headers_policy" {
				name        = "test-consume-modify-headers-policy"
				description = "Test consume policy for modifying headers"
				condition   = "context.topic.name == 'header-topic'"
				enabled     = true

				config = {
					actions = [
						{
							set = {
								key   = "X-Custom-Header"
								value = "custom-value"
							}
						},
						{
							remove = {
								key = "X-Remove-Header"
							}
						}
					]
				}
			}
		`)
		require.NoError(t, err)

		consumePolicyModifyHeaders.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)
		consumePolicyModifyHeaders.AddAttribute(
			"virtual_cluster_id",
			egwVirtualCluster.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{

				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicyModifyHeaders).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
								plancheck.ResourceActionCreate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"name",
							"test-consume-modify-headers-policy",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"description",
							"Test consume policy for modifying headers",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"enabled",
							"true",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"condition",
							"context.topic.name == 'header-topic'",
						),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(
							consumePolicyModifyHeaders.AddAttribute("labels", `{ env = "dev", service = "headers" }`),
						).Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"labels.%",
							"2",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"labels.env",
							"dev",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"labels.service",
							"headers",
						),
					),
				},
			},
		})
	})

	t.Run("EGW Consume Policy Schema Validation", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)

		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		consumePolicySchemaValidation, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_schema_validation" "test_schema_validation_policy" {
				name        = "test-consume-schema-validation-policy"
				description = "Test consume policy for schema validation"
				condition   = "context.topic.name == 'validated-topic'"
				enabled     = true

				config = {
					type                     = "json"
					key_validation_action    = "mark"
					value_validation_action  = "skip"
				}
			}
		`)
		require.NoError(t, err)

		consumePolicySchemaValidation.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)
		consumePolicySchemaValidation.AddAttribute(
			"virtual_cluster_id",
			egwVirtualCluster.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySchemaValidation).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "name", "test-consume-schema-validation-policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "description", "Test consume policy for schema validation"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "enabled", "true")),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySchemaValidation).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySchemaValidation.AddAttribute("labels", `{ env = "test", validation = "strict" }`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "labels.%", "2"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "labels.env", "test"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "labels.validation", "strict"),
					),
				},
			},
		})
	})

	t.Run("EGW Consume Policy Skip Record", func(t *testing.T) {

		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)

		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		consumePolicySkipRecord, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_skip_record" "test_skip_record_policy" {
				name        = "test_skip_record_policy"
				description = "Test consume policy for skipping records"
				condition   = "context.topic.name == 'skip-topic'"
				enabled     = true
			}
		`)
		require.NoError(t, err)

		consumePolicySkipRecord.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)
		consumePolicySkipRecord.AddAttribute(
			"virtual_cluster_id",
			egwVirtualCluster.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySkipRecord).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "name", "test_skip_record_policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "description", "Test consume policy for skipping records"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "condition", "context.topic.name == 'skip-topic'")),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySkipRecord).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySkipRecord.AddAttribute("condition", `context.topic.name == 'new-skip-topic'`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "condition", "context.topic.name == 'new-skip-topic'"),
					),
				},
			},
		})
	})

	t.Run("EGW Data Plane Certificate", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egw, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		cert, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_data_plane_certificate" "tf-test-dp-cert" {
			  name        = "tf-test-dp-cert"
			  description = "initial certificate"
    certificate = <<EOF
-----BEGIN CERTIFICATE-----
MIIB4TCCAYugAwIBAgIUAenxUyPjkSLCe2BQXoBMBacqgLowDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yNDEwMjgyMDA3NDlaFw0zNDEw
MjYyMDA3NDlaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwXDANBgkqhkiG9w0BAQEF
AANLADBIAkEAyzipjrbAaLO/yPg7lL1dLWzhqNdc3S4YNR7f1RG9whWhbsPE2z42
e6WGFf9hggP6xjG4qbU8jFVczpd1UPwGbQIDAQABo1MwUTAdBgNVHQ4EFgQUkPPB
ghj+iHOHAKJlC1gLbKT/ZHQwHwYDVR0jBBgwFoAUkPPBghj+iHOHAKJlC1gLbKT/
ZHQwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAANBALfy49GvA2ld+u+G
Koxa8kCt7uywoqu0hfbBfUT4HqmXPvsuhz8RinE5ltxId108vtDNlD/+bKl+N5Ub
qKjBs0k=
-----END CERTIFICATE-----
EOF
			}
		`)
		require.NoError(t, err)

		cert.AddAttribute(
			"gateway_id",
			egw.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egw).Upsert(cert).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_data_plane_certificate.tf-test-dp-cert",
								plancheck.ResourceActionCreate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_data_plane_certificate.tf-test-dp-cert",
							"name",
							"tf-test-dp-cert",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_data_plane_certificate.tf-test-dp-cert",
							"description",
							"initial certificate",
						),
					),
				},
				{
					Config: builder.Upsert(egw).Upsert(cert).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.
						Upsert(egw).Upsert(cert.AddAttribute("name", `"tf-test-dp-cert-updated"`).
						AddAttribute("description", `"updated certificate description"`),
					).Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_data_plane_certificate.tf-test-dp-cert",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_data_plane_certificate.tf-test-dp-cert",
							"name",
							"tf-test-dp-cert-updated",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_data_plane_certificate.tf-test-dp-cert",
							"description",
							"updated certificate description",
						),
					),
				},
			},
		})
	})

	t.Run("EGW Listener", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect
		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)
		egwListener, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener" "test_listener" {
				name        = "test-listener"
				description = "Test listener configuration"
				addresses   = ["0.0.0.0"]
				ports       = ["9092"]
			}
		`)
		require.NoError(t, err)
		egwListener.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwListener).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_listener.test_listener",
								plancheck.ResourceActionCreate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"name",
							"test-listener",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"description",
							"Test listener configuration",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"addresses.0",
							"0.0.0.0",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"ports.0",
							"9092",
						),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwListener).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwListener.AddAttribute("description", `Test listener configuration`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"description",
							"Test listener configuration",
						),
					),
				},
			},
		})
	})

	t.Run("EGW Listener Policy Forward To Virtual Cluster", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)

		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwListener, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener" "tf_test_egw_listener" {
				name      = "tf_test_egw_listener"
				addresses = ["0.0.0.0"]
				ports     = ["9092", "9093", "9094", "9095"]
				labels    = {}

				lifecycle {
					ignore_changes = [labels]
				}
				description = "Description for listener"
			}
		`)
		require.NoError(t, err)

		egwListener.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		listenerPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener_policy_forward_to_virtual_cluster" "tf_test_egw_forward_policy" {
				name        = "tf_test_egw_forward_policy"
				description = "Test forward to virtual cluster policy"
				enabled     = true

				config = {
					port_mapping = {
						advertised_host = "broker.example.com"
						bootstrap_port  = "at_start"
						min_broker_id   = 0
						destination = {
							id = konnect_event_gateway_virtual_cluster.my_event_gateway_virtual_cluster.id
						}
					}
				}
			}
		`)
		require.NoError(t, err)

		listenerPolicy.AddAttribute(
			"listener_id",
			egwListener.ResourcePath()+".id",
		)
		listenerPolicy.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(egwListener).Upsert(listenerPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "name", "tf_test_egw_forward_policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "description", "Test forward to virtual cluster policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "enabled", "true")),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(listenerPolicy.AddAttribute("description", "Update description")).Build(),
					Check:  resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "description", "Update description")),
				},
			},
		})
	})

	t.Run("EGW Listener Policy Forward To Virtual Cluster with SNI", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)

		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwListener, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener" "tf_test_egw_listener" {
				name      = "tf_test_egw_listener"
				addresses = ["0.0.0.0"]
				ports     = ["9092", "9093", "9094", "9095"]
				labels    = {}

				lifecycle {
					ignore_changes = [labels]
				}
				description = "Description for listener"
			}
		`)
		require.NoError(t, err)

		egwListener.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		listenerPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener_policy_forward_to_virtual_cluster" "tf_test_egw_forward_policy" {
				name        = "tf_test_egw_forward_policy"
				description = "Test forward to virtual cluster policy"
				enabled     = true

				config = {
					sni = {
						sni_suffix = ".example.com"
					}
				}
			}
		`)
		require.NoError(t, err)

		listenerPolicy.AddAttribute(
			"listener_id",
			egwListener.ResourcePath()+".id",
		)
		listenerPolicy.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(egwListener).Upsert(listenerPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "name", "tf_test_egw_forward_policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "description", "Test forward to virtual cluster policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "config.sni.sni_suffix", ".example.com"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(egwListener).Upsert(listenerPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(egwListener).Upsert(listenerPolicy.AddAttribute("description", `"Updated description"`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "description", "Updated description"),
					),
				},
			},
		})
	})

	t.Run("EGW Listener Policy TLS Server", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwListener, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener" "test_listener_tls" {
				name      = "test-listener-tls"
				addresses = ["0.0.0.0"]
				ports     = ["9092", "9093"]
				labels    = {}
				description = "Listener for TLS server policy"

				lifecycle {
					ignore_changes = [labels]
				}
			}
		`)
		require.NoError(t, err)

		egwListener.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		tlsServerPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener_policy_tls_server" "test_tls_policy" {
				name        = "test-tls-server-policy"
				description = "Test TLS server policy for encryption"
				enabled     = true

				config = {
					allow_plaintext = false
					certificates = [
						{
							certificate = "-----BEGIN CERTIFICATE-----\nMIIBxjCCAUygAwIBAgIUX9TaLbWF76yQc8IGR+YRbeiDlHkwCgYIKoZIzj0EAwIwGjEYMBYGA1UEAwwPa29uZ19jbHVzdGVyaW5nMB4XDTI0MDMwMTE0MzkxNloXDTI3MDMwMTE0MzkxNlowGjEYMBYGA1UEAwwPa29uZ19jbHVzdGVyaW5nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEcMndCotXzeZ9vGAMfDfZ7UxUuP5bcIrwwUOI8YlpMdvB12HvjtS7O0/ONr3fBeCWagRuitPEqd4b3EJuD8kuFUMt+2A09N6KY1YDJWgKHei7rzKgrefzVt11XgBiDsUBo1MwUTAdBgNVHQ4EFgQUIrdAC8p02h60GZW0Jlh2Vcg/WeMwHwYDVR0jBBgwFoAUIrdAC8p02h60GZW0Jlh2Vcg/WeMwDwYDVR0TAQH/BAUwAwEB/zAKBggqhkjOPQQDAgNoADBlAjBYb+yQf33sItlmsONLc41Agtx73FMEN7LfWA85OtlkMie1N1x0mj08pzS/Xc1VONwCMQDN9sBn3Kody0gse+EXYSuPPj1oo9jmFB9/xrpz35YpDATvuyhH8xwSJ4xMuxQiduc=\n-----END CERTIFICATE-----"
							key = "-----BEGIN PRIVATE KEY-----\nMIG2AgEAMBAGByqGSM49AgEGBSuBBAAiBIGeMIGbAgEBBDDLuRX+uzSbstvLWsQrWwuGK4AdjLU/tN9A/fn03gxNvppKw++SBtnLyB+9YZ29YA+hZANiAARwyd0Ki1fN5n28YAx8N9ntTFS4/ltwivDBQ4jxiWkx28HXYe+O1Ls7T842vd8F4JZqBG6K08Sp3hvcQm4PyS4VQy37YDT03opjVgMlaAod6LuvMqCt5/NW3XVeAGIOxQE=\n-----END PRIVATE KEY-----"
						}
					]
					versions = {
						min = "TLSv1.2"
						max = "TLSv1.3"
					}
				}
			  }
		`)
		require.NoError(t, err)

		require.NoError(t, err)

		tlsServerPolicy.AddAttribute(
			"listener_id",
			egwListener.ResourcePath()+".id",
		)
		tlsServerPolicy.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwListener).Upsert(tlsServerPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", "name", "test-tls-server-policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", "description", "Test TLS server policy for encryption"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", "config.allow_plaintext", "false")),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwListener).Upsert(tlsServerPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwListener.AddAttribute("description", "Update description")).Build(),
					Check:  resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttr("konnect_event_gateway_listener.test_listener_tls", "description", "Update description")),
				},
			},
		})
	})

	t.Run("EGW Produce Policy Encrypt", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
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

	t.Run("EGW Produce Policy Modify Headers", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		backendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		backendCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		virtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)

		virtualCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		policyCreate, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_modify_headers" "test-produce-headers-policy" {
			  name        = "test-produce-headers-policy"
			  description = "Test produce policy for modifying headers"
			  condition   = "context.topic.name == 'header-topic'"
			  enabled     = true

			  config = {
				actions = [
				  {
					set = {
					  key   = "X-Producer-Header"
					  value = "producer-value"
					}
				  },
				  {
					remove = {
					  key = "X-Remove-Header"
					}
				  }
				]
			  }

			  gateway_id         = konnect_event_gateway.my_event_gateway.id
			  virtual_cluster_id = konnect_event_gateway_virtual_cluster.my_event_gateway_virtual_cluster.id
			}
		`)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(backendCluster).Upsert(virtualCluster).Upsert(policyCreate).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "name", "test-produce-headers-policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "description", "Test produce policy for modifying headers"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "condition", "context.topic.name == 'header-topic'"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "enabled", "true")),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(backendCluster).Upsert(virtualCluster).Upsert(policyCreate).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(backendCluster).Upsert(virtualCluster).Upsert(policyCreate.AddAttribute("description", "Updated Test produce policy for modifying headers")).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "description", "Updated Test produce policy for modifying headers")),
				},
			},
		})
	})

	t.Run("EGW Static Key", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		staticKey, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_static_key" "test_egw_static_key" {
			  name  = "test_egw_static_key"
			  value = "$${vault.env['MY_ENV_VAR']}"
			}
		`)
		require.NoError(t, err)

		staticKey.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{

				{
					Config: builder.
						Upsert(egwCp).
						Upsert(staticKey).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_static_key.test_egw_static_key", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_static_key.test_egw_static_key", "name", "test_egw_static_key"),
						resource.TestCheckResourceAttrSet("konnect_event_gateway_static_key.test_egw_static_key", "id"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(staticKey).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(staticKey.AddAttribute("labels", `{key = "value"}`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_static_key.test_egw_static_key", "labels.%", "1"),
						resource.TestCheckResourceAttr("konnect_event_gateway_static_key.test_egw_static_key", "labels.key", "value"),
					),
				},
			},
		})
	})

	t.Run("EGW Cluster Policy ACLs", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egw, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		backend, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)

		backend.AddAttribute(
			"gateway_id",
			egw.ResourcePath()+".id",
		)

		virtualCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster" "tf-test-virtual-acls" {
			  name      = "tf-test-virtual-acls"
			  dns_label = "tf-test-vc-acls"
			  acl_mode  = "enforce_on_gateway"

			  authentication = [
				{
				  anonymous = {}
				}
			  ]

			  destination = {
				id = konnect_event_gateway_backend_cluster.my_event_gateway_backend_cluster.id
			  }
			}
		`)
		require.NoError(t, err)

		virtualCluster.AddAttribute(
			"gateway_id",
			egw.ResourcePath()+".id",
		)

		clusterPolicyACLs, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_cluster_policy_acls" "tf-test-cluster-policy-acls" {
			  name        = "tf-test-cluster-policy-acls"
			  description = "initial description"
			  enabled     = false
			  condition   = "context.topic.name.endsWith('my_suffix')"

			  labels = {
				key = "value"
			  }

			  config = {
				rules = [
				  {
					action = "deny"

					operations = [
					  {
						name = "describe_configs"
					  }
					]

					resource_names = [
					  {
						match = "...my_match..."
					  }
					]

					resource_type = "transactional_id"
				  }
				]
			  }
			}
		`)
		require.NoError(t, err)

		clusterPolicyACLs.AddAttribute(
			"gateway_id",
			egw.ResourcePath()+".id",
		)
		clusterPolicyACLs.AddAttribute(
			"virtual_cluster_id",
			virtualCluster.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egw).Upsert(backend).Upsert(virtualCluster).Upsert(clusterPolicyACLs).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_cluster_policy_acls.tf-test-cluster-policy-acls",
								plancheck.ResourceActionCreate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_cluster_policy_acls.tf-test-cluster-policy-acls",
							"enabled",
							"false",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_cluster_policy_acls.tf-test-cluster-policy-acls",
							"config.rules.0.action",
							"deny",
						),
					),
				},
				{
					Config: builder.Upsert(egw).Upsert(backend).Upsert(virtualCluster).Upsert(clusterPolicyACLs).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},

				{
					Config: builder.Upsert(egw).Upsert(backend).Upsert(virtualCluster).Upsert(
						clusterPolicyACLs.AddAttribute("enabled", "true")).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_cluster_policy_acls.tf-test-cluster-policy-acls",
								plancheck.ResourceActionUpdate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_cluster_policy_acls.tf-test-cluster-policy-acls",
							"enabled",
							"true",
						),
					),
				},
			},
		})
	})

	t.Run("EGW Produce Policy Schema Validation", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		backendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)
		backendCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		virtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)
		virtualCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		schemaRegistry, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_schema_registry" "tf_test_schema_registry" {
			  confluent = {
				name        = "my-schema-registry"
				description = "schema registry for produce policy"

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

		schemaRegistry.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		producePolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_schema_validation" "tf_test_produce_schema_validation" {
			  name        = "tf-test-produce-schema-validation"
			  description = "initial schema validation policy"
			  enabled     = true

			  config = {
				confluent_schema_registry = {
				  key_validation_action   = "reject"
				  value_validation_action = "reject"
				  schema_registry = {
					  id = konnect_event_gateway_schema_registry.tf_test_schema_registry.id
				  }
				}
			  }
			}
		`)
		require.NoError(t, err)

		producePolicy.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		producePolicy.AddAttribute("virtual_cluster_id", virtualCluster.ResourcePath()+".id")

		producePolicy.AddAttribute(
			"depends_on",
			`[konnect_event_gateway_schema_registry.tf_test_schema_registry]`,
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(backendCluster).Upsert(virtualCluster).Upsert(schemaRegistry).
						Upsert(producePolicy).Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_produce_policy_schema_validation.tf_test_produce_schema_validation",
								plancheck.ResourceActionCreate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_produce_policy_schema_validation.tf_test_produce_schema_validation",
							"enabled",
							"true",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_produce_policy_schema_validation.tf_test_produce_schema_validation",
							"config.confluent_schema_registry.key_validation_action",
							"reject",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_produce_policy_schema_validation.tf_test_produce_schema_validation",
							"config.confluent_schema_registry.value_validation_action",
							"reject",
						),
					),
				},

				{
					Config: builder.Upsert(egwCp).Upsert(backendCluster).Upsert(virtualCluster).Upsert(schemaRegistry).
						Upsert(producePolicy).Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},

				{
					Config: builder.Upsert(egwCp).Upsert(backendCluster).Upsert(virtualCluster).Upsert(schemaRegistry).
						Upsert(
							producePolicy.
								AddAttribute("enabled", "false").
								AddAttribute("description", `"updated schema validation policy"`),
						).Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_produce_policy_schema_validation.tf_test_produce_schema_validation",
								plancheck.ResourceActionUpdate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_produce_policy_schema_validation.tf_test_produce_schema_validation",
							"enabled",
							"false",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_produce_policy_schema_validation.tf_test_produce_schema_validation",
							"description",
							"updated schema validation policy",
						),
					),
				},
			},
		})
	})

	t.Run("EGW Confluent Schema Registry", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		schemaRegistry, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_schema_registry" "test_schema_registry" {
			  confluent = {
				name        = "test-schema-registry"
				description = "test schema registry"

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

		schemaRegistry.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(schemaRegistry).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_schema_registry.test_schema_registry",
								plancheck.ResourceActionCreate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_schema_registry.test_schema_registry",
							"confluent.name",
							"test-schema-registry",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_schema_registry.test_schema_registry",
							"confluent.config.schema_type",
							"avro",
						),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(schemaRegistry).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(
						schemaRegistry.AddAttribute(
							"confluent.labels",
							`{ env = "test" }`,
						),
					).Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_schema_registry.test_schema_registry",
							"confluent.labels.%", "1",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_schema_registry.test_schema_registry",
							"confluent.labels.env",
							"test",
						),
					),
				},
			},
		})
	})

}
