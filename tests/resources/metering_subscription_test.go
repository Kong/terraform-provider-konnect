package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMeteringSubscription(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		// Subscription destroy is unclear - and customer, plan deletion gets blocked.
		t.Skip("Subscription destroy is unclear - customer/plan deletion gets blocked")
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					// Create with the required customer/plan references plus one optional field.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrPair("konnect_metering_subscription.my_meteringsubscription", "customer.id", "konnect_metering_customer.my_meteringcustomer", "id"),
						resource.TestCheckResourceAttrPair("konnect_metering_subscription.my_meteringsubscription", "plan.id", "konnect_metering_plan.my_meteringplan", "id"),
					),
				},
				{
					// Update plan
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrPair("konnect_metering_subscription.my_meteringsubscription", "plan.id", "konnect_metering_plan.my_meteringplan2", "id"),
					),
				},
			},
		})
	})
}
