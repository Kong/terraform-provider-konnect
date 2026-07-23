package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMeshControlPlanes(t *testing.T) {
	t.Parallel()

	t.Run("data source with filter", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"data.konnect_mesh_control_planes.my_meshcontrolplanes",
							tfjsonpath.New("data").AtSliceIndex(0).AtMapKey("name"),
							knownvalue.StringExact("Lookup Mesh Control Plane"),
						),
					},
				},
			},
		})
	})
}
