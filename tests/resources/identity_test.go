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
	identityAuthServer = `
		resource "konnect_identity_auth_server" "my_authserver" {
		  name          = "my-auth-server"
		  audience      = "local-demo"
			force_destroy = true
		}
	`

	identityAuthServerScope = `
		resource "konnect_identity_auth_server_scope" "my_scope" {
		  name                = "my-scope"
		  description         = "My Scope"
		  include_in_metadata = true
		  enabled             = true

		  auth_server_id = konnect_identity_auth_server.my_authserver.id
		}
	`

	identityAuthServerClaim = `
		resource "konnect_identity_auth_server_claim" "my_authserverclaims" {
		  enabled               = true
		  include_in_all_scopes = false
		  include_in_scopes = [
		    konnect_identity_auth_server_scope.my_scope.id
		  ]
		  include_in_token = true
		  name             = "my-claim"
		  value            = "some-value"

		  auth_server_id = konnect_identity_auth_server.my_authserver.id
		}
	`

	identityAuthServerClient = `
		resource "konnect_identity_auth_server_client" "my_client" {
		  name             = "my-client"
		  allow_all_scopes = true
		  grant_types = [
		    "client_credentials",
		    "implicit"
		  ]

		  response_types = [
		    "id_token",
		    "token",
		    "code"
		  ]

		  auth_server_id = konnect_identity_auth_server.my_authserver.id
		}
	`
)

func TestIdentityResources(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("Auth Server", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		authServer, err := hclbuilder.FromString(identityAuthServer)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(authServer).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_identity_auth_server.my_authserver", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_identity_auth_server.my_authserver", "name", "my-auth-server"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server.my_authserver", "audience", "local-demo"),
						resource.TestCheckResourceAttrSet("konnect_identity_auth_server.my_authserver", "issuer"),
						resource.TestCheckResourceAttrSet("konnect_identity_auth_server.my_authserver", "metadata_uri"),
					),
				},
				{
					Config: builder.Upsert(authServer).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("Auth Server Scope", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		authServer, err := hclbuilder.FromString(identityAuthServer)
		require.NoError(t, err)

		scope, err := hclbuilder.FromString(identityAuthServerScope)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(authServer).Upsert(scope).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_identity_auth_server_scope.my_scope", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_identity_auth_server_scope.my_scope", "name", "my-scope"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_scope.my_scope", "description", "My Scope"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_scope.my_scope", "include_in_metadata", "true"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_scope.my_scope", "enabled", "true"),
					),
				},
				{
					Config: builder.Upsert(authServer).Upsert(scope).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("Auth Server Claim", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		authServer, err := hclbuilder.FromString(identityAuthServer)
		require.NoError(t, err)

		scope, err := hclbuilder.FromString(identityAuthServerScope)
		require.NoError(t, err)

		claim, err := hclbuilder.FromString(identityAuthServerClaim)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(authServer).Upsert(scope).Upsert(claim).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_identity_auth_server_claim.my_authserverclaims", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_identity_auth_server_claim.my_authserverclaims", "name", "my-claim"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_claim.my_authserverclaims", "value", "some-value"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_claim.my_authserverclaims", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_claim.my_authserverclaims", "include_in_all_scopes", "false"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_claim.my_authserverclaims", "include_in_token", "true"),
						resource.TestCheckResourceAttrSet("konnect_identity_auth_server_claim.my_authserverclaims", "include_in_scopes.0"),
					),
				},
				{
					Config: builder.Upsert(authServer).Upsert(scope).Upsert(claim).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("Auth Server Client", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		authServer, err := hclbuilder.FromString(identityAuthServer)
		require.NoError(t, err)

		client, err := hclbuilder.FromString(identityAuthServerClient)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(authServer).Upsert(client).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_identity_auth_server_client.my_client", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_identity_auth_server_client.my_client", "name", "my-client"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_client.my_client", "allow_all_scopes", "true"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_client.my_client", "response_types.0", "id_token"),
						resource.TestCheckResourceAttrSet("konnect_identity_auth_server_client.my_client", "id"),
						resource.TestCheckResourceAttrSet("konnect_identity_auth_server_client.my_client", "client_secret"),
					),
				},
				{
					Config: builder.Upsert(authServer).Upsert(client).Build(),
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
