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
	catalogMcp = `
		resource "konnect_catalog_mcp" "test_mcp" {
			name         = "test-mcp-tf"
			display_name = "Test MCP"
			description  = "MCP for testing"
			labels = {
				env = "test"
			}
		}
	`

	catalogMcpVersion = `
		resource "konnect_catalog_mcp_version" "test_mcp_version" {
			mcp_id  = konnect_catalog_mcp.test_mcp.id
			version = "1.0.0"
			resources = [
				{
					name        = "config-file"
					uri         = "file:///etc/mcp/config.json"
					title       = "MCP Configuration"
					description = "Configuration file for the MCP"
					mime_type   = "application/json"
					size        = 1234
				}
			]
		}
	`

	testAiGatewayForMcp = `
		resource "konnect_ai_gateway" "test_ai_gateway_mcp" {
			provider     = konnect-beta
			name         = "test-ai-gateway-mcp-impl"
			display_name = "Test AI Gateway for MCP"
		}
	`

	testMcpServer = `
		resource "konnect_ai_gateway_mcp_server" "test_mcp_server" {
			provider   = konnect-beta
			gateway_id = konnect_ai_gateway.test_ai_gateway_mcp.id
			upstream_server = {
				name         = "test-mcp-server"
				display_name = "Test MCP Server"
				config = {
					url = "https://mcp.example.com"
					tools_cache_ttl_seconds = 300
					route = {
						paths = ["/mcp"]
						hosts = []
					}
				}
			}
		}
	`

	catalogMcpImplementation = `
		resource "konnect_catalog_mcp_implementation" "test_mcp_impl" {
			mcp_id = konnect_catalog_mcp.test_mcp.id
			create_catalog_mcp_gateway_implementation = {
				implementation = {
					config = {
						gateway_control_plane_id = konnect_ai_gateway.test_ai_gateway_mcp.id
						gateway_mcp_server_id    = konnect_ai_gateway_mcp_server.test_mcp_server.id
					}
					type = "ai-gateway"
				}
			}
		}
	`
)

func TestCatalogMCP(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("Catalog MCP", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		mcp, err := hclbuilder.FromString(catalogMcp)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(mcp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_catalog_mcp.test_mcp", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_catalog_mcp.test_mcp", "name", "test-mcp-tf"),
						resource.TestCheckResourceAttr("konnect_catalog_mcp.test_mcp", "display_name", "Test MCP"),
					),
				},
				{
					Config: builder.Upsert(mcp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(mcp.AddAttribute("description", `"Updated MCP for testing"`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_catalog_mcp.test_mcp", "description", "Updated MCP for testing"),
					),
				},
			},
		})
	})

	t.Run("Catalog MCP Version", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		mcp, err := hclbuilder.FromString(catalogMcp)
		require.NoError(t, err)

		version, err := hclbuilder.FromString(catalogMcpVersion)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(mcp).Upsert(version).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_catalog_mcp_version.test_mcp_version", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_catalog_mcp_version.test_mcp_version", "id"),
						resource.TestCheckResourceAttr("konnect_catalog_mcp_version.test_mcp_version", "version", "1.0.0")),
				},
				{
					Config: builder.Upsert(mcp).Upsert(version).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(mcp).Upsert(version.AddAttribute("version", `"2.0.0"`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_catalog_mcp_version.test_mcp_version", "version", "2.0.0"),
					),
				},
			},
		})
	})

	t.Run("Catalog MCP Implementation", func(t *testing.T) {
		builder := hclbuilder.New()

		mcp, err := hclbuilder.FromString(catalogMcp)
		require.NoError(t, err)

		gateway, err := hclbuilder.FromString(testAiGatewayForMcp)
		require.NoError(t, err)

		server, err := hclbuilder.FromString(testMcpServer)
		require.NoError(t, err)

		impl, err := hclbuilder.FromString(catalogMcpImplementation)
		require.NoError(t, err)

		baseConfig := builder.Upsert(mcp).Upsert(gateway).Upsert(server).Upsert(impl).Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			ExternalProviders: map[string]resource.ExternalProvider{
				"konnect-beta": {
					Source:            "Kong/konnect-beta",
					VersionConstraint: "0.22.0",
				},
			},
			Steps: []resource.TestStep{
				{
					Config: baseConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_catalog_mcp_implementation.test_mcp_impl", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_catalog_mcp_implementation.test_mcp_impl", "id"),
						resource.TestCheckResourceAttrSet("konnect_catalog_mcp_implementation.test_mcp_impl", "mcp_id")),
				},
				{
					Config: baseConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_catalog_mcp_implementation.test_mcp_impl", plancheck.ResourceActionNoop),
						},
					},
				},
			},
		})
	})
}
