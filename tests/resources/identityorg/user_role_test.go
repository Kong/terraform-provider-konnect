package identityorg

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUserRole(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("konnect_user_role.my_userrole", "entity_id"),
						resource.TestCheckResourceAttr("konnect_user_role.my_userrole", "role_name", "Viewer"),
					),
				},
			},
		})
	})

	t.Run("entity-ids", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_user_role.my_userrole", "entity_ids.#", "1"),
						resource.TestCheckResourceAttrPair(
							"konnect_user_role.my_userrole", "entity_ids.0",
							"konnect_gateway_control_plane.my_konnect_cp", "id",
						),
						resource.TestCheckResourceAttr("konnect_user_role.my_userrole", "role_name", "Viewer"),
					),
				},
			},
		})
	})

	// Rejected at plan time: both ids set.
	t.Run("mutually-exclusive", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ExpectError:     regexp.MustCompile(`Attribute "entity_ids" cannot be specified when "entity_id" is\s+specified`),
				},
			},
		})
	})
}
