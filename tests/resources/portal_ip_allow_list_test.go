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
	portalResource = `
		resource "konnect_portal" "test_portal" {
		  name                     = "test_portal_for_ip_allow_list"
		  authentication_enabled   = true
		  auto_approve_applications = false
		  auto_approve_developers   = false
		  default_api_visibility    = "public"
		  default_page_visibility   = "private"
		  force_destroy             = "true"
		  rbac_enabled              = false
		}
	`
	portalIPAllowListResource = `
		resource "konnect_portal_ip_allow_list" "test_ip_allow_list" {
		  allowed_ips = [
		    "192.168.1.1",
		    "192.168.1.0/22"
		  ]
		}
	`
)

func TestPortalIPAllowList(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("Portal IP Allow List Create Update Delete", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		portal, err := hclbuilder.FromString(portalResource)
		require.NoError(t, err)

		ipAllowList, err := hclbuilder.FromString(portalIPAllowListResource)
		require.NoError(t, err)

		ipAllowList.AddAttribute(
			"portal_id",
			portal.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				// Step 1: Create
				{
					Config: builder.Upsert(portal).Upsert(ipAllowList).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_portal_ip_allow_list.test_ip_allow_list", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_ip_allow_list.test_ip_allow_list", "allowed_ips.#", "2"),
						resource.TestCheckResourceAttr("konnect_portal_ip_allow_list.test_ip_allow_list", "allowed_ips.0", "192.168.1.1"),
						resource.TestCheckResourceAttr("konnect_portal_ip_allow_list.test_ip_allow_list", "allowed_ips.1", "192.168.1.0/22"),
						resource.TestCheckResourceAttrSet("konnect_portal_ip_allow_list.test_ip_allow_list", "id"),
						resource.TestCheckResourceAttrSet("konnect_portal_ip_allow_list.test_ip_allow_list", "portal_id"),
					),
				},
				{
					Config: builder.Upsert(portal).Upsert(ipAllowList).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(portal).Upsert(
						ipAllowList.AddAttribute("allowed_ips", `["192.168.1.1", "192.168.1.0/22", "10.0.0.1"]`),
					).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_portal_ip_allow_list.test_ip_allow_list", plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_ip_allow_list.test_ip_allow_list", "allowed_ips.#", "3"),
						resource.TestCheckResourceAttr("konnect_portal_ip_allow_list.test_ip_allow_list", "allowed_ips.0", "192.168.1.1"),
						resource.TestCheckResourceAttr("konnect_portal_ip_allow_list.test_ip_allow_list", "allowed_ips.1", "192.168.1.0/22"),
						resource.TestCheckResourceAttr("konnect_portal_ip_allow_list.test_ip_allow_list", "allowed_ips.2", "10.0.0.1"),
					),
				},
			},
		})
	})
}
