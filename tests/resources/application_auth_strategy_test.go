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
		    key_names     = ["apikey"]
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
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_applicationauthstrategy", "key_auth.key_names.0", "apikey"),
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_applicationauthstrategy", "key_auth.configs.key_auth.key_names.0", "apikey"),
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
				{
					Config: builder.Upsert(authStrategy.AddAttribute("key_auth.display_name", `"Updated Test Strategy"`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_application_auth_strategy.my_applicationauthstrategy", plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_application_auth_strategy.my_applicationauthstrategy", "key_auth.display_name", "Updated Test Strategy"),
					),
				},
			},
		})
	})
}
