package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	testAuthServerForSecret = `
resource "konnect_identity_auth_server" "my_authserver" {
  name          = "tf-ci-testing-authserver-for-client-secret"
  audience      = "local-demo"
  force_destroy = true
}
`

	testAuthServerClientForSecret = `
resource "konnect_identity_auth_server_client" "my_client" {
  name             = "tf-ci-testing-client-for-secret"
  allow_all_scopes = true
  grant_types = [
    "client_credentials"
  ]

  response_types = [
    "code"
  ]

  auth_server_id = konnect_identity_auth_server.my_authserver.id
}
`
)

func TestIdentityAuthServerClientSecret(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()

	t.Run("should do CRUD", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		authServer, err := hclbuilder.FromString(testAuthServerForSecret)
		require.NoError(t, err)

		authServerClient, err := hclbuilder.FromString(testAuthServerClientForSecret)
		require.NoError(t, err)

		clientSecret, err := hclbuilder.FromString(`resource "konnect_identity_auth_server_client_secret" "my_client_secret" {
			auth_server_id = konnect_identity_auth_server.my_authserver.id
			client_id      = konnect_identity_auth_server_client.my_client.id
			enabled        = true
			secret         = "YAzsyUlNZ5gNGeKS9H3VAdxVPzhPo4ae"
			}
			`)
		require.NoError(t, err)

		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(authServer).Upsert(authServerClient).Upsert(clientSecret).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_identity_auth_server.my_authserver", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_identity_auth_server_client.my_client", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_identity_auth_server_client_secret.my_client_secret", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_identity_auth_server_client_secret.my_client_secret", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_client_secret.my_client_secret", "secret", "YAzsyUlNZ5gNGeKS9H3VAdxVPzhPo4ae"),
						resource.TestCheckResourceAttrSet("konnect_identity_auth_server_client_secret.my_client_secret", "id"),
						resource.TestCheckResourceAttrSet("konnect_identity_auth_server_client_secret.my_client_secret", "created_at"),
						resource.TestCheckResourceAttrSet("konnect_identity_auth_server_client_secret.my_client_secret", "updated_at"),
					),
				},
				// Update: only `enabled` is mutable in place; toggling it must not replace the secret.
				{
					Config: builder.Upsert(authServer).Upsert(authServerClient).Upsert(clientSecret.AddAttribute("enabled", "false")).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_identity_auth_server_client_secret.my_client_secret", plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_identity_auth_server_client_secret.my_client_secret", "enabled", "false"),
						resource.TestCheckResourceAttr("konnect_identity_auth_server_client_secret.my_client_secret", "secret", "YAzsyUlNZ5gNGeKS9H3VAdxVPzhPo4ae"),
					),
				},
			},
		})
	})
}
