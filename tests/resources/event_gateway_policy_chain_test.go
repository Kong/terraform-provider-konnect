package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestEventGatewayPolicyChain(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	// Shared resources for all tests
	const eventGatewayCP = `
		resource "konnect_event_gateway" "test_egw" {
		  name = "test_policy_chain_egw"
		}
	`

	const eventGatewayBackendCluster = `
		resource "konnect_event_gateway_backend_cluster" "test_backend_cluster" {
		  name = "test_policy_chain_backend_cluster"
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

	const eventGatewayVirtualCluster = `
		resource "konnect_event_gateway_virtual_cluster" "test_virtual_cluster" {
		  name      = "test_policy_chain_virtual_cluster"
		  dns_label = "vcluster-policy-chain"
		  acl_mode  = "enforce_on_gateway"
		  authentication = [
		    {
		      anonymous = {}
		    }
		  ]
		  destination = {
		    id = konnect_event_gateway_backend_cluster.test_backend_cluster.id
		  }
		}
	`

	t.Run("Listener Policy Chain", func(t *testing.T) {
		t.Skip("Skipping listener policy chain test until supported in the Provider")
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)
		egwBackendCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)
		egwVirtualCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		egwListener, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener" "test_listener" {
				name        = "test-listener-policy-chain"
				description = "Listener for policy chain test"
				addresses   = ["0.0.0.0"]
				ports       = ["9092", "9093", "9094", "9095"]
			}
		`)
		require.NoError(t, err)
		egwListener.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		listenerPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener_policy_forward_to_virtual_cluster" "test_listener_policy" {
				name        = "test-listener-forward-policy"
				description = "Test forward policy for chain"
				enabled     = true

				config = {
					port_mapping = {
						advertised_host = "broker.example.com"
						bootstrap_port  = "at_start"
						min_broker_id   = 0
						destination = {
							id = konnect_event_gateway_virtual_cluster.test_virtual_cluster.id
						}
					}
				}
			}
		`)
		require.NoError(t, err)
		listenerPolicy.AddAttribute("listener_id", egwListener.ResourcePath()+".id")
		listenerPolicy.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		listenerPolicyTLS, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener_policy_tls_server" "test_listener_policy_tls" {
				name        = "test-listener-tls-policy"
				description = "Test TLS policy for chain"
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
		listenerPolicyTLS.AddAttribute("listener_id", egwListener.ResourcePath()+".id")
		listenerPolicyTLS.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		listenerPolicyChain, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener_policy_chain" "test_listener_policy_chain" {
				policies = [
					{
						id = konnect_event_gateway_listener_policy_forward_to_virtual_cluster.test_listener_policy.id
					}
				]
			}
		`)
		require.NoError(t, err)
		listenerPolicyChain.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		listenerPolicyChain.AddAttribute("listener_id", egwListener.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				// Create
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(egwListener).Upsert(listenerPolicy).Upsert(listenerPolicyChain).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_listener_policy_chain.test_listener_policy_chain",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener_policy_chain.test_listener_policy_chain",
							"policies.#",
							"1",
						),
					),
				},
				// Verify no changes on re-apply
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(egwListener).Upsert(listenerPolicy).Upsert(listenerPolicyChain).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				// Update - add second policy to chain
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(egwListener).Upsert(listenerPolicy).Upsert(listenerPolicyTLS).Upsert(listenerPolicyChain.AddAttribute("policies", `[
							{
								id = konnect_event_gateway_listener_policy_tls_server.test_listener_policy_tls.id
							},
							{
								id = konnect_event_gateway_listener_policy_forward_to_virtual_cluster.test_listener_policy.id
							}
						]`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_listener_policy_chain.test_listener_policy_chain",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener_policy_chain.test_listener_policy_chain",
							"policies.#",
							"2",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_event_gateway_listener_policy_chain.test_listener_policy_chain",
							"policies.0.id",
							"konnect_event_gateway_listener_policy_tls_server.test_listener_policy_tls",
							"id",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_event_gateway_listener_policy_chain.test_listener_policy_chain",
							"policies.1.id",
							"konnect_event_gateway_listener_policy_forward_to_virtual_cluster.test_listener_policy",
							"id",
						),
					),
				},
			},
		})
	})

	t.Run("Cluster Policy Chain", func(t *testing.T) {
		t.Skip("Skipping cluster policy chain test until supported in the Provider")
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)
		egwBackendCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)
		egwVirtualCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		clusterPolicyACL, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_cluster_policy_acls" "test_cluster_policy_acl" {
				name        = "test-cluster-acl-policy"
				description = "Test ACL policy for chain"
				enabled     = true
				condition   = "context.topic.name == 'test-topic'"

				config = {
					rules = [
						{
							action = "deny"
							resource_type = "topic"
							operations = [
							{
								name = "describe"
							}
							]
							resource_names = [
							{
								match = "topic_*"
							}
							]
						}
					]
				}
			}
		`)
		require.NoError(t, err)
		clusterPolicyACL.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		clusterPolicyACL.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")

		clusterPolicyChain, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster_cluster_policy_chain" "test_cluster_policy_chain" {
				policies = [
					{
						id = konnect_event_gateway_cluster_policy_acls.test_cluster_policy_acl.id
					}
				]
			}
		`)
		require.NoError(t, err)
		clusterPolicyChain.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		clusterPolicyChain.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")

		clusterPolicyACL2, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_cluster_policy_acls" "test_cluster_policy_acl_2" {
				name        = "test-cluster-acl-policy-2"
				description = "Second ACL policy for chain update"
				enabled     = true
				condition   = "context.topic.name == 'another-topic'"

				config = {
					rules = [
						{
							action = "deny"
							resource_type = "topic"
							operations = [
							{
								name = "describe"
							}
							]
							resource_names = [
							{
								match = "topic_*"
							}
							]
						}
					]
				}
			}
		`)
		require.NoError(t, err)
		clusterPolicyACL2.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		clusterPolicyACL2.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				// Create
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(clusterPolicyACL).Upsert(clusterPolicyChain).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_virtual_cluster_cluster_policy_chain.test_cluster_policy_chain",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster_cluster_policy_chain.test_cluster_policy_chain",
							"policies.#",
							"1",
						),
					),
				},
				// Verify no changes on re-apply
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(clusterPolicyACL).Upsert(clusterPolicyChain).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				// Update - add second policy to chain
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(clusterPolicyACL).Upsert(clusterPolicyACL2).Upsert(clusterPolicyChain.AddAttribute("policies", `[
							{
								id = konnect_event_gateway_cluster_policy_acls.test_cluster_policy_acl_2.id
							},
							{
								id = konnect_event_gateway_cluster_policy_acls.test_cluster_policy_acl.id
							}
						]`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_virtual_cluster_cluster_policy_chain.test_cluster_policy_chain",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster_cluster_policy_chain.test_cluster_policy_chain",
							"policies.#",
							"2",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_event_gateway_virtual_cluster_cluster_policy_chain.test_cluster_policy_chain",
							"policies.0.id",
							"konnect_event_gateway_cluster_policy_acls.test_cluster_policy_acl_2",
							"id",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_event_gateway_virtual_cluster_cluster_policy_chain.test_cluster_policy_chain",
							"policies.1.id",
							"konnect_event_gateway_cluster_policy_acls.test_cluster_policy_acl",
							"id",
						),
					),
				},
			},
		})
	})

	t.Run("Produce Policy Chain", func(t *testing.T) {
		t.Skip("Skipping produce policy chain test until supported in the Provider")
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)
		egwBackendCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)
		egwVirtualCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		producePolicyModifyHeaders, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_modify_headers" "test_produce_policy" {
				name        = "test-produce-modify-headers"
				description = "Test produce policy for chain"
				condition   = "context.topic.name == 'produce-topic'"
				enabled     = true

				config = {
					actions = [
						{
							set = {
								key   = "X-Producer-Header"
								value = "producer-value"
							}
						}
					]
				}
			}
		`)
		require.NoError(t, err)
		producePolicyModifyHeaders.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		producePolicyModifyHeaders.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")

		producePolicyChain, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster_produce_policy_chain" "test_produce_policy_chain" {
				policies = [
					{
						id = konnect_event_gateway_produce_policy_modify_headers.test_produce_policy.id
					}
				]
			}
		`)
		require.NoError(t, err)
		producePolicyChain.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		producePolicyChain.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")

		producePolicyModifyHeaders2, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_modify_headers" "test_produce_policy_2" {
				name        = "test-produce-modify-headers-2"
				description = "Second produce policy for chain update"
				condition   = "context.topic.name == 'another-produce-topic'"
				enabled     = true

				config = {
					actions = [
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
		producePolicyModifyHeaders2.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		producePolicyModifyHeaders2.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				// Create
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(producePolicyModifyHeaders).Upsert(producePolicyChain).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_virtual_cluster_produce_policy_chain.test_produce_policy_chain",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster_produce_policy_chain.test_produce_policy_chain",
							"policies.#",
							"1",
						),
					),
				},
				// Verify no changes on re-apply
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(producePolicyModifyHeaders).Upsert(producePolicyChain).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				// Update - add second policy to chain
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(producePolicyModifyHeaders).Upsert(producePolicyModifyHeaders2).Upsert(producePolicyChain.AddAttribute("policies", `[
							{
								id = konnect_event_gateway_produce_policy_modify_headers.test_produce_policy_2.id
							},
							{
								id = konnect_event_gateway_produce_policy_modify_headers.test_produce_policy.id
							}
						]`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_virtual_cluster_produce_policy_chain.test_produce_policy_chain",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster_produce_policy_chain.test_produce_policy_chain",
							"policies.#",
							"2",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_event_gateway_virtual_cluster_produce_policy_chain.test_produce_policy_chain",
							"policies.0.id",
							"konnect_event_gateway_produce_policy_modify_headers.test_produce_policy_2",
							"id",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_event_gateway_virtual_cluster_produce_policy_chain.test_produce_policy_chain",
							"policies.1.id",
							"konnect_event_gateway_produce_policy_modify_headers.test_produce_policy",
							"id",
						),
					),
				},
			},
		})
	})

	t.Run("Consume Policy Chain", func(t *testing.T) {
		t.Skip("Skipping consume policy chain test until supported in the Provider")
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		egwCp, err := hclbuilder.FromString(eventGatewayCP)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(eventGatewayBackendCluster)
		require.NoError(t, err)
		egwBackendCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		egwVirtualCluster, err := hclbuilder.FromString(eventGatewayVirtualCluster)
		require.NoError(t, err)
		egwVirtualCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		consumePolicySkipRecord, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_skip_record" "test_consume_policy" {
				name        = "test-consume-skip-record"
				description = "Test consume policy for chain"
				condition   = "context.topic.name == 'consume-topic'"
				enabled     = true
			}
		`)
		require.NoError(t, err)
		consumePolicySkipRecord.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		consumePolicySkipRecord.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")

		consumePolicyChain, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster_consume_policy_chain" "test_consume_policy_chain" {
				policies = [
					{
						id = konnect_event_gateway_consume_policy_skip_record.test_consume_policy.id
					}
				]
			}
		`)
		require.NoError(t, err)
		consumePolicyChain.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		consumePolicyChain.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")

		// Create a second consume policy for update test (modify headers)
		consumePolicyModifyHeaders, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_modify_headers" "test_consume_policy_2" {
				name        = "test-consume-modify-headers"
				description = "Second consume policy for chain update"
				condition   = "context.topic.name == 'another-consume-topic'"
				enabled     = true

				config = {
					actions = [
						{
							set = {
								key   = "X-Consumer-Header"
								value = "consumer-value"
							}
						}
					]
				}
			}
		`)
		require.NoError(t, err)
		consumePolicyModifyHeaders.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")
		consumePolicyModifyHeaders.AddAttribute("virtual_cluster_id", egwVirtualCluster.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				// Create
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(consumePolicySkipRecord).Upsert(consumePolicyChain).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_virtual_cluster_consume_policy_chain.test_consume_policy_chain",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster_consume_policy_chain.test_consume_policy_chain",
							"policies.#",
							"1",
						),
					),
				},
				// Verify no changes on re-apply
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(consumePolicySkipRecord).Upsert(consumePolicyChain).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				// Update - add second policy to chain
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).
						Upsert(consumePolicySkipRecord).Upsert(consumePolicyModifyHeaders).Upsert(consumePolicyChain.AddAttribute("policies", `[
							{
								id = konnect_event_gateway_consume_policy_modify_headers.test_consume_policy_2.id
							},
							{
								id = konnect_event_gateway_consume_policy_skip_record.test_consume_policy.id
							}
						]`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_virtual_cluster_consume_policy_chain.test_consume_policy_chain",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster_consume_policy_chain.test_consume_policy_chain",
							"policies.#",
							"2",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_event_gateway_virtual_cluster_consume_policy_chain.test_consume_policy_chain",
							"policies.0.id",
							"konnect_event_gateway_consume_policy_modify_headers.test_consume_policy_2",
							"id",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_event_gateway_virtual_cluster_consume_policy_chain.test_consume_policy_chain",
							"policies.1.id",
							"konnect_event_gateway_consume_policy_skip_record.test_consume_policy",
							"id",
						),
					),
				},
			},
		})
	})
}
