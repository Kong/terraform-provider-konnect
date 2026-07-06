package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMeteringMeter(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// Create with required fields plus one optional field
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_metering_meter.my_meteringmeter", "aggregation", "count"),
						resource.TestCheckResourceAttr("konnect_metering_meter.my_meteringmeter", "event_type", "prompt"),
						resource.TestCheckResourceAttr("konnect_metering_meter.my_meteringmeter", "key", "tftestmeterkey"),
						resource.TestCheckResourceAttr("konnect_metering_meter.my_meteringmeter", "name", "tftestmetername"),
						resource.TestCheckResourceAttr("konnect_metering_meter.my_meteringmeter", "description", "aaaa"),
					),
				},
				{
					// Update description field
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_metering_meter.my_meteringmeter", "description", "bbbb"),
					),
				},
			},
		})
	})
}
