package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMeteringFeature(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// Create with all required fields plus one optional field.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_metering_feature.my_meteringfeature", "key", "tftestfeaturekey"),
						resource.TestCheckResourceAttr("konnect_metering_feature.my_meteringfeature", "name", "tftestfeaturename"),
						resource.TestCheckResourceAttr("konnect_metering_feature.my_meteringfeature", "unit_cost.manual.amount", "1.00"),
					),
				},
				{
					// Update the optional unit_cost field in place.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_metering_feature.my_meteringfeature", "unit_cost.manual.amount", "2.00"),
					),
				},
			},
		})
	})
}
