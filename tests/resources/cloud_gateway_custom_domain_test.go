package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestCloudGatewayCustomDomain(t *testing.T) {
	t.Run("Cloud Gateway Custom Domain owned by API Gateway", func(t *testing.T) {
		suffix := acctest.RandStringFromCharSet(8, acctest.CharSetAlpha)
		domain := fmt.Sprintf("tf-test-api-%s.example.com", suffix)

		builder := hclbuilder.New()
		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "test_cp" {
             name          = "tf-test-cp-custom-domain"
             cloud_gateway = true
          }
       `)
		require.NoError(t, err)
		customDomain, err := hclbuilder.FromString(fmt.Sprintf(`
          resource "konnect_cloud_gateway_custom_domain" "my_custom_domain" {
             control_plane_geo = "us"
             control_plane_id  = konnect_gateway_control_plane.test_cp.id
             domain            = "%s"
          }
       `, domain))
		require.NoError(t, err)

		fullConfig := builder.
			Upsert(cp).
			Upsert(customDomain).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: fullConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_cloud_gateway_custom_domain.my_custom_domain",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_cloud_gateway_custom_domain.my_custom_domain", "domain", domain),

						//validating if type  defaults to API.
						resource.TestCheckResourceAttr("konnect_cloud_gateway_custom_domain.my_custom_domain", "type", "api"),
						resource.TestCheckResourceAttrPair(
							"konnect_cloud_gateway_custom_domain.my_custom_domain", "control_plane_id",
							"konnect_gateway_control_plane.test_cp", "id",
						),
					),
				},
				{
					Config: fullConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("Cloud Gateway Custom Domain owned by AI Gateway", func(t *testing.T) {
		suffix := acctest.RandStringFromCharSet(8, acctest.CharSetAlpha)
		domain := fmt.Sprintf("tf-test-ai-%s.example.com", suffix)

		builder := hclbuilder.New()
		aiGateway, err := hclbuilder.FromString(`
          resource "konnect_ai_gateway" "test_ai_gateway" {
             name            = "tf-test-ai-gateway-custom-domain"
             display_name    = "TF Test AI Gateway Custom Domain"
             deployment_type = "managed"
          }
       `)
		require.NoError(t, err)
		customDomain, err := hclbuilder.FromString(fmt.Sprintf(`
          resource "konnect_cloud_gateway_custom_domain" "my_ai_custom_domain" {
             control_plane_geo = "us"
             control_plane_id  = konnect_ai_gateway.test_ai_gateway.id
             domain            = "%s"
             kind              = "dedicated.v0"
             type              = "ai"
          }
       `, domain))
		require.NoError(t, err)

		fullConfig := builder.
			Upsert(aiGateway).
			Upsert(customDomain).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: fullConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_cloud_gateway_custom_domain.my_ai_custom_domain",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_cloud_gateway_custom_domain.my_ai_custom_domain", "domain", domain),
						resource.TestCheckResourceAttr("konnect_cloud_gateway_custom_domain.my_ai_custom_domain", "type", "ai"),
					),
				},
				{
					Config: fullConfig,
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
