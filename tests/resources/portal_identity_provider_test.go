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
	portalForIdentityProvider = `
		resource "konnect_portal" "my_portal" {
			force_destroy = true
			name          = "TFAcceptancePortalIdentityProvider"
		}
	`
	portalIdentityProviderOIDC = `
		resource "konnect_portal_identity_provider" "oidc_provider" {
			type      = "oidc"
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
)

func TestPortalIdentityProvider(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("oidc", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		portal, err := hclbuilder.FromString(portalForIdentityProvider)
		require.NoError(t, err)

		oidcProvider, err := hclbuilder.FromString(portalIdentityProviderOIDC)
		require.NoError(t, err)
		oidcProvider.AddAttribute("portal_id", portal.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(portal).Upsert(oidcProvider).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_portal_identity_provider.oidc_provider", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_identity_provider.oidc_provider", "type", "oidc"),
						resource.TestCheckResourceAttrSet("konnect_portal_identity_provider.oidc_provider", "id"),
						resource.TestCheckResourceAttrSet("konnect_portal_identity_provider.oidc_provider", "portal_id"),
					),
				},
				{
					Config: builder.Upsert(portal).Upsert(oidcProvider.AddAttribute("enabled", "true")).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_portal_identity_provider.oidc_provider", "type", "oidc"),
						resource.TestCheckResourceAttr("konnect_portal_identity_provider.oidc_provider", "enabled", "true"),
						resource.TestCheckResourceAttrSet("konnect_portal_identity_provider.oidc_provider", "id"),
						resource.TestCheckResourceAttrSet("konnect_portal_identity_provider.oidc_provider", "portal_id"),
					),
				},

				{
					Config: builder.Upsert(portal).Upsert(oidcProvider).Build(),
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
