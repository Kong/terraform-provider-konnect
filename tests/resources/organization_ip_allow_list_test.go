package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestOrgIPAllowList(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			Steps: []resource.TestStep{
				{
					ProtoV6ProviderFactories: providerFactory,
					Config:                   providerConfigUs,
					ConfigDirectory:          config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_organization_ip_allow_list.my_organizationipallowlist", "enabled", "false"),
						resource.TestCheckResourceAttr("konnect_organization_ip_allow_list.my_organizationipallowlist", "allowed_ips.#", "1"),
					),
				},
				{
					ProtoV6ProviderFactories: providerFactory,
					Config:                   providerConfigUs,
					ConfigDirectory:          config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_organization_ip_allow_list.my_organizationipallowlist", "allowed_ips.#", "2"),
					),
				},
				{
					// Regression check for a perpetual diff: the API is free to
					// return allowed_ips in a different order than it was sent,
					// and this must not show up as a plan-time change (this is
					// exactly the bug that was found and fixed by switching the
					// attribute to a Set instead of a List).
					ProtoV6ProviderFactories: providerFactory,
					Config:                   providerConfigUs,
					ConfigDirectory:          config.TestStepDirectory(),
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
