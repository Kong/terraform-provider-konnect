package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestSystemAccount(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_system_account.my_systemaccount", "name", "TFAcceptanceSystemAccountName"),
						resource.TestCheckResourceAttr("konnect_system_account.my_systemaccount", "description", "my description"),
					),
				},
				{
					// Update some fields
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_system_account.my_systemaccount", "description", "my updated description"),
					),
				},
			},
		})
	})

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
						resource.TestCheckOutput("system_account_id", "75b799f9-316e-4fa8-b62f-b8a238ee266d"),
					),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownOutputValueAtPath(
							"system_accounts",
							tfjsonpath.New("data").AtSliceIndex(0).AtMapKey("name"),
							knownvalue.StringExact("Sample System Account"),
						),
					},
				},
			},
		})
	})
}
