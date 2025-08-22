package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAPIDocument(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_api_document.my_parent_apidocument", "title", "Parent API Document"),
						resource.TestCheckResourceAttr("konnect_api_document.my_child_apidocument", "title", "Child API Document"),
						resource.TestCheckResourceAttr("konnect_api_document.my_child_apidocument", "status", "unpublished"),
						resource.TestCheckResourceAttrSet("konnect_api_document.my_child_apidocument", "parent_document_id"),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_api_document.my_child_apidocument", "status", "published"),
						resource.TestCheckNoResourceAttr("konnect_api_document.my_child_apidocument", "parent_document_id"),
					),
				},
			},
		})
	})
}
