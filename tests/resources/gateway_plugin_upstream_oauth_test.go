package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayPluginUpstreamOauth(t *testing.T) {
	t.Parallel()

	// OIDC scopes must be optional, and default to an empty array
	t.Run("empty scopes", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:             providerConfigUs,
					ConfigDirectory:    config.TestNameDirectory(),
					ExpectNonEmptyPlan: true, // todo: remove after false diff is fixed
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_plugin_upstream_oauth.test", "config.oauth.scopes.#", "0"),
					),
				},
			},
		})
	})
}
