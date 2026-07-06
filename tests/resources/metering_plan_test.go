package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMeteringPlan(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// Create with all required fields plus one optional field.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_metering_plan.my_meteringplan", "billing_cadence", "P1Y"),
						resource.TestCheckResourceAttr("konnect_metering_plan.my_meteringplan", "currency", "USD"),
						resource.TestCheckResourceAttr("konnect_metering_plan.my_meteringplan", "key", "resource_keyy"),
						resource.TestCheckResourceAttr("konnect_metering_plan.my_meteringplan", "name", "Myfirstplan"),
						resource.TestCheckResourceAttr("konnect_metering_plan.my_meteringplan", "description", "aaaa"),
					),
				},
				{
					// Update the optional description field in place.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_metering_plan.my_meteringplan", "description", "bbbb"),
					),
				},
			},
		})
	})
}
