package tests

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGatewayCredentials(t *testing.T) {
	t.Run("fail_if_consumer_set", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ExpectError:     regexp.MustCompile("An argument named \"consumer\" is not expected here"),
				},
			},
		})
	})
}
