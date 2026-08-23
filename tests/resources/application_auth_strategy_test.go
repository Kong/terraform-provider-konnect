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
	applicationAuthStrategy = `
		resource "konnect_application_auth_strategy" "my_applicationauthstrategy" {
		  key_auth = {
		    name          = "my-application-auth-strategy"
		    display_name  = "My Test Strategy"
		    strategy_type = "key_auth"
		    configs = {
		      key_auth = {
		        key_names = ["apikey"]
		      }
		    }
		  }
		}
	`

	applicationAuthStrategyOpenIDConnect = `
    resource "konnect_application_auth_strategy" "my_oidc_applicationauthstrategy" {
      openid_connect = {
        name          = "my-oidc-application-auth-strategy"
        display_name  = "My OIDC Test Strategy"
        strategy_type = "openid_connect"
        configs = {
          openid_connect = {
            issuer            = "https://issuer.example.com"
            credential_claim  = ["sub"]
            scopes            = ["openid"]
            auth_methods      = ["authorization_code"]
          }
        }
      }
    }`
)

func TestApplicationAuthStrategy(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("Application Auth Strategy", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		authStrategy, err := hclbuilder.FromString(applicationAuthStrategy)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(authStrategy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_application_auth_strategy.my_applicationauthstrategy", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_applicationauthstrategy", "key_auth.name", "my-application-auth-strategy"),
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_applicationauthstrategy", "key_auth.display_name", "My Test Strategy"),
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_applicationauthstrategy", "key_auth.strategy_type", "key_auth"),
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_applicationauthstrategy", "key_auth.configs.key_auth.key_names.0", "apikey"),
					),
				},
			},
		})
	})
}

// TestApplicationAuthStrategyUpdate verifies that changing key_names on an existing
// key_auth strategy updates the resource in place
func TestApplicationAuthStrategyUpdate(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("updating key_names updates in place", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		authStrategy, err := hclbuilder.FromString(applicationAuthStrategy)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(authStrategy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_application_auth_strategy.my_applicationauthstrategy", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_applicationauthstrategy", "key_auth.configs.key_auth.key_names.0", "apikey"),
					),
				},
				{
					Config: builder.Upsert(
						authStrategy.AddAttribute("key_auth.configs.key_auth.key_names", `["apikey", "x-api-key"]`),
					).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_application_auth_strategy.my_applicationauthstrategy", plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_applicationauthstrategy", "key_auth.configs.key_auth.key_names.0", "apikey"),
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_applicationauthstrategy", "key_auth.configs.key_auth.key_names.1", "x-api-key"),
					),
				},
				{
					Config: builder.Upsert(authStrategy).Build(),
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

// TestApplicationAuthStrategyOpenIDConnectUpdate verifies that changing scopes on an
// existing openid_connect strategy updates the resource in place
func TestApplicationAuthStrategyOpenIDConnectUpdate(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("updating scopes updates in place", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		authStrategy, err := hclbuilder.FromString(applicationAuthStrategyOpenIDConnect)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(authStrategy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_application_auth_strategy.my_oidc_applicationauthstrategy", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_oidc_applicationauthstrategy", "openid_connect.configs.openid_connect.scopes.0", "openid"),
					),
				},
				{
					Config: builder.Upsert(
						authStrategy.AddAttribute("openid_connect.configs.openid_connect.scopes", `["openid", "profile"]`),
					).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_application_auth_strategy.my_oidc_applicationauthstrategy", plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_oidc_applicationauthstrategy", "openid_connect.configs.openid_connect.scopes.0", "openid"),
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_oidc_applicationauthstrategy", "openid_connect.configs.openid_connect.scopes.1", "profile"),
					),
				},
				{
					Config: builder.Upsert(authStrategy).Build(),
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
