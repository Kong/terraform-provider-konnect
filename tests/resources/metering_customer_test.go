package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMeteringCustomer(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_metering_customer.my_meteringcustomer", "key", "tftestcustomerkey"),
						resource.TestCheckResourceAttr("konnect_metering_customer.my_meteringcustomer", "name", "tftestcustomername"),
						resource.TestCheckResourceAttr("konnect_metering_customer.my_meteringcustomer", "description", "aaaa"),
					),
				},
				{
					// Update description field
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_metering_customer.my_meteringcustomer", "description", "bbbb"),
					),
				},
			},
		})
	})
}
