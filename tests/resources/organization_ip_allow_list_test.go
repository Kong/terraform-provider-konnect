package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// The allow list is always created with enabled = false so that running these
// tests can never lock the test organization (or CI) out of the Admin API.
func TestOrgIPAllowList(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_organization_ip_allow_list.my_organizationipallowlist", "enabled", "false"),
						resource.TestCheckResourceAttr("konnect_organization_ip_allow_list.my_organizationipallowlist", "allowed_ips.#", "1"),
					),
				},
				{
					// Adds an address and lists them in reverse order: allowed_ips is a
					// set, so the order the API returns must not produce a diff.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_organization_ip_allow_list.my_organizationipallowlist", "allowed_ips.#", "2"),
					),
				},
			},
		})
	})
}
