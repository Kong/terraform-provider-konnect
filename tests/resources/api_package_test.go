package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestAPIPackageReusable(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("API Package", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		apiPackage, err := hclbuilder.FromString(`
			resource "konnect_api_package" "my_api_package" {
			  name        = "my-test-api-package"
			  description = "My API Package Description"
			  version     = "v1"
			  slug        = "my-test-api-package-v1"

			  rate_limiting_config = {
			    duration = "hours"
			    limit    = 1000
			  }
			}
		`)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(apiPackage).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_api_package.my_api_package", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_api_package.my_api_package", "id"),
						resource.TestCheckResourceAttr("konnect_api_package.my_api_package", "name", "my-test-api-package"),
						resource.TestCheckResourceAttr("konnect_api_package.my_api_package", "slug", "my-test-api-package-v1"),
						resource.TestCheckResourceAttr("konnect_api_package.my_api_package", "rate_limiting_config.duration", "hours"),
					),
				},
				{
					Config: builder.Upsert(apiPackage).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(apiPackage.AddAttribute("labels", `{env = "staging"}`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_api_package.my_api_package", "labels.%", "1"),
						resource.TestCheckResourceAttr("konnect_api_package.my_api_package", "labels.env", "staging"),
					),
				},
			},
		})
	})
}
