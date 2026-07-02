package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestMeteringCustomer(t *testing.T) {
	t.Run("plan-diff", func(t *testing.T) {
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

	t.Run("update-nullify-fields", func(t *testing.T) {
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
					// Update some fields to null
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckNoResourceAttr("konnect_metering_customer.my_meteringcustomer", "description"),
					),
				},
			},
		})
	})
}
