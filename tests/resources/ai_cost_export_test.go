package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

// The range is relative to the run because the API rejects a range that ends
// after the day in progress. Seconds precision keeps the value the API echoes
// back identical to the configured one.
const aiCostExport = `
	resource "konnect_ai_cost_export" "test_export" {
		compression = "gzip"
		dataset     = "usage_events"
		format      = "csv"
		from        = "%s"
		to          = "%s"
		time_zone   = "UTC"

		// A record keyed by usage field, not the array a budget scope takes.
		filters = jsonencode({
			provider = { eq = "openai" }
		})

		labels = {
			reason = "acceptance-test"
		}
	}
`

func TestAICostExport(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()

	now := time.Now().UTC().Truncate(time.Second)
	from := now.Add(-48 * time.Hour).Format(time.RFC3339)
	to := now.Add(-24 * time.Hour).Format(time.RFC3339)

	newConfig := func(t *testing.T) string {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		export, err := hclbuilder.FromString(fmt.Sprintf(aiCostExport, from, to))
		require.NoError(t, err)
		builder.Upsert(export)

		return builder.Build()
	}

	t.Run("Usage Events Export", func(t *testing.T) {
		const address = "konnect_ai_cost_export.test_export"

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: newConfig(t),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(address, "id"),
						resource.TestCheckResourceAttr(address, "dataset", "usage_events"),
						resource.TestCheckResourceAttr(address, "format", "csv"),
						resource.TestCheckResourceAttr(address, "compression", "gzip"),
						resource.TestCheckResourceAttr(address, "time_zone", "UTC"),
						resource.TestCheckResourceAttr(address, "from", from),
						resource.TestCheckResourceAttr(address, "to", to),
						resource.TestCheckResourceAttr(address, "labels.reason", "acceptance-test"),
						resource.TestCheckResourceAttrSet(address, "status"),
					),
				},
				{
					// Every attribute requires replacement, so an unchanged config has
					// nothing to plan.
					Config: newConfig(t),
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
