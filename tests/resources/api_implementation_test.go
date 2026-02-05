package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAPIImplementation(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_api_implementation.my_apiimplementation", "api_id"),
						resource.TestCheckResourceAttrSet("konnect_api_implementation.my_apiimplementation", "service_reference.service.id"),
						resource.TestCheckResourceAttrSet("konnect_api_implementation.my_apiimplementation", "service_reference.service.control_plane_id"),
						resource.TestCheckResourceAttrSet("konnect_api_implementation.my_apiimplementation_cp_linked", "api_id"),
						resource.TestCheckResourceAttrSet("konnect_api_implementation.my_apiimplementation_cp_linked", "control_plane_reference.control_plane.id"),
					),
				},
			},
		})
	})
}
