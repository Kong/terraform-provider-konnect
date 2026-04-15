package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestCloudGatewayNetwork(t *testing.T) {
	t.Run("data source with name filter", func(t *testing.T) {
		t.Skip("CGW tests can not be run in CI due to long-running resource creation/deletion. This test has been run manually")
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:             providerConfigUs,
					ConfigDirectory:    config.TestNameDirectory(),
					ExpectNonEmptyPlan: false,
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownOutputValueAtPath(
							"network_name",
							tfjsonpath.New("name"),
							knownvalue.StringExact("ci-test"),
						),
					},
				},
			},
		})
	})
}
