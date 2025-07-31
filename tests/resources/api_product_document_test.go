package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAPIProductDocument(t *testing.T) {
	t.Run("plan-diff", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_api_product_document.my_apiproductdocument", "title", "My Konnect API Product Document"),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	/*
		// Todo: Re-enable this test once the changes to nullify are made, otherwise leads to dangling resources.
		t.Run("update-nullify-fields", func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: providerFactory,
				Steps: []resource.TestStep{
					{
						Config:          providerConfigUs,
						ConfigDirectory: config.TestNameDirectory(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_api_product_document.my_apiproduct_child_document", "slug", "path-for-seo-child"),
						),
					},
					{
						// Update some fields to null
						Config:          providerConfigUs,
						ConfigDirectory: config.TestStepDirectory(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckNoResourceAttr("konnect_api_product_document.my_apiproduct_child_document", "parent_document_id"),
						),
					},
				},
			})
		})
	*/
}
