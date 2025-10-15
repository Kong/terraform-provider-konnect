package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestGatewayControlPlane(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName    string
		name        string
		clusterType string
	}{
		{
			testName:    "hybrid",
			name:        "Terraform Control Plane",
			clusterType: "CLUSTER_TYPE_CONTROL_PLANE",
		},
		{
			testName:    "ingress controller",
			name:        "KIC Terraform Control Plane",
			clusterType: "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER",
		},
	}

	for _, tc := range tests {
		t.Run("type "+tc.testName, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: providerFactory,
				Steps: []resource.TestStep{
					{
						Config:          providerConfigUs,
						ConfigDirectory: config.TestNameDirectory(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("konnect_gateway_control_plane.tfdemo", "name", tc.name),
							resource.TestCheckResourceAttr("konnect_gateway_control_plane.tfdemo", "cluster_type", tc.clusterType),
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
	}

	t.Run("requires replacement", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_gateway_control_plane.tfdemo", plancheck.ResourceActionDestroyBeforeCreate),
						},
					},
				},
			},
		})
	})

	t.Run("data source", func(t *testing.T) {

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:             providerConfigUs,
					ConfigDirectory:    config.TestNameDirectory(),
					ExpectNonEmptyPlan: true, // outputs become known
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckOutput("control_plane", "Lookup Control Plane"),
					),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownOutputValueAtPath(
							"control_plane_list",
							tfjsonpath.New("data").AtSliceIndex(0).AtMapKey("name"),
							knownvalue.StringExact("Lookup Control Plane"),
						),
					},
				},
			},
		})
	})
}
