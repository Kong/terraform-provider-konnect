package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestPlatformAvailableRegions(t *testing.T) {
	t.Run("data", func(t *testing.T) {

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						// Verify that the nested attributes exist in state
						resource.TestCheckResourceAttrSet("data.konnect_platform_available_regions.test", "regions.stable.#"),
						resource.TestCheckResourceAttrSet("data.konnect_platform_available_regions.test", "regions.beta.#"),
						resource.TestCheckResourceAttrSet("data.konnect_platform_available_regions.test", "regions.stable_opt_in.#"),
					),
				},
			},
		})
	})
}
