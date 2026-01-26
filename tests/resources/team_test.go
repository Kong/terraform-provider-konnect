package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestTeam(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_team.my_team", "name", "TFAcceptanceTeamName"),
					),
				},
				{
					// Update description field
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_team.my_team", "description", "TF acceptance test team updated description."),
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
						resource.TestCheckOutput("team_id", "0d591155-3e7f-4352-992d-eaea72fdf7b1"),
					),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownOutputValueAtPath(
							"teams",
							tfjsonpath.New("data").AtSliceIndex(0).AtMapKey("name"),
							knownvalue.StringExact("portal-admin"),
						),
					},
				},
			},
		})
	})
}
