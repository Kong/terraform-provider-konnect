package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/require"
)

const addressOfCostManagementTags = "konnect_ai_gateway_cost_management_tags.my_aigatewaycostmanagementtags"

const (
	aiGatewayCostManagementTags = `
		resource "konnect_ai_gateway_cost_management_tags" "my_aigatewaycostmanagementtags" {
			headers = [
				{
					name    = "tftest-tenant-id"
					tag_key = "tftest-tenant"
				}
			]
			prefixes = [
				{
					name = "tftest-co"
				}
			]
		}
	`

	aiGatewayCostManagementTagsUpdated = `
		resource "konnect_ai_gateway_cost_management_tags" "my_aigatewaycostmanagementtags" {
			headers = [
				{
					name    = "tftest-tenant-id"
					tag_key = "tftest-tenant-updated"
				}
			]
			prefixes = [
				{
					name = "tftest-co"
				}
			]
		}
	`
)


func TestAIGatewayCostManagementTags(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()

	newConfig := func(t *testing.T, tags string) string {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		tagsResource, err := hclbuilder.FromString(tags)
		require.NoError(t, err)
		builder.Upsert(tagsResource)

		return builder.Build()
	}

	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: newConfig(t, aiGatewayCostManagementTags),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(addressOfCostManagementTags, "headers.0.tag_key", "tftest-tenant"),
						resource.TestCheckResourceAttr(addressOfCostManagementTags, "prefixes.0.name", "tftest-co"),
					),
				},
				{
					// Update headers.0.tag_key
					Config: newConfig(t, aiGatewayCostManagementTagsUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(addressOfCostManagementTags, "headers.0.tag_key", "tftest-tenant-updated"),
					),
				},
			},
		})
	})
}
