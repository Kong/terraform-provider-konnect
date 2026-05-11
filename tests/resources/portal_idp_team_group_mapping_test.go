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
	portalForIdpMapping = `
		resource "konnect_portal" "my_portal" {
			force_destroy = true
			name          = "TFAcceptancePortalIdpMapping"
		}
	`
	portalTeamForIdpMapping = `
		resource "konnect_portal_team" "my_team" {
			name        = "TFAcceptancePortalTeamGroupMapping"
			description = "TF acceptance test portal team for group mapping."
		}
	`
	portalIdpOIDCForMapping = `
		resource "konnect_portal_identity_provider" "oidc_provider" {
			type = "oidc"
			config = {
				oidc_identity_provider_config = {
					client_id     = "client-id"
					client_secret = "my-client-secret"
					issuer_url    = "https://accounts.google.com"
					scopes = [
						"openid",
						"email",
						"profile"
					]
					claim_mappings = {
						email  = "email"
						groups = "groups"
						name   = "name"
					}
				}
			}
		}
	`

	portalIdpTeamGroupMappingOIDC = `
		resource "konnect_portal_idp_team_group_mapping" "my_mapping" {
			group = "Tech Leads"
		}
	`
)

func TestPortalIdpTeamGroupMapping(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("oidc", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		portal, err := hclbuilder.FromString(portalForIdpMapping)
		require.NoError(t, err)

		portalTeam, err := hclbuilder.FromString(portalTeamForIdpMapping)
		require.NoError(t, err)
		portalTeam.AddAttribute("portal_id", portal.ResourcePath()+".id")

		oidcProvider, err := hclbuilder.FromString(portalIdpOIDCForMapping)
		require.NoError(t, err)
		oidcProvider.AddAttribute("portal_id", portal.ResourcePath()+".id")

		mapping, err := hclbuilder.FromString(portalIdpTeamGroupMappingOIDC)
		require.NoError(t, err)
		mapping.AddAttribute("id", oidcProvider.ResourcePath()+".id")
		mapping.AddAttribute("portal_id", portal.ResourcePath()+".id")
		mapping.AddAttribute("team_id", portalTeam.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(portal).Upsert(portalTeam).Upsert(oidcProvider).Upsert(mapping).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_portal_idp_team_group_mapping.my_mapping", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_idp_team_group_mapping.my_mapping", "group", "Tech Leads"),
						resource.TestCheckResourceAttrSet("konnect_portal_idp_team_group_mapping.my_mapping", "id"),
						resource.TestCheckResourceAttrSet("konnect_portal_idp_team_group_mapping.my_mapping", "portal_id"),
						resource.TestCheckResourceAttrSet("konnect_portal_idp_team_group_mapping.my_mapping", "team_id"),
						resource.TestCheckResourceAttrSet("konnect_portal_idp_team_group_mapping.my_mapping", "mapping_id"),
					),
				},
				{
					Config: builder.Upsert(portal).Upsert(portalTeam).Upsert(oidcProvider).Upsert(mapping).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

}
