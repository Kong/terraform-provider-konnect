package aigateway

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

const (
	aiCostExportSchedule = `
		resource "konnect_ai_cost_export_schedule" "test_schedule" {
			name        = "tf-test-daily-usage"
			cadence     = "daily"
			dataset     = "usage_events"
			format      = "csv"
			compression = "gzip"
			status      = "active"
			time_zone   = "UTC"

			// A record keyed by usage field, not the array a budget scope takes.
			filters = jsonencode({
				provider = { in = ["openai", "anthropic"] }
			})

			labels = {
				team = "finops"
			}
		}
	`

	aiCostExportScheduleUpdated = `
		resource "konnect_ai_cost_export_schedule" "test_schedule" {
			name        = "tf-test-daily-usage-paused"
			cadence     = "daily"
			dataset     = "usage_events"
			format      = "csv"
			compression = "gzip"
			status      = "paused"
			time_zone   = "UTC"

			labels = {
				team = "finops"
			}
		}
	`
	// A change to `cadence` forces a replacement: the API refuses it on an existing
	// schedule. Without the `x-speakeasy-param-force-new` annotation the provider
	// would send an update and the API would answer 400.
	aiCostExportScheduleWeekly = `
		resource "konnect_ai_cost_export_schedule" "test_schedule" {
			name        = "tf-test-daily-usage-paused"
			cadence     = "weekly"
			dataset     = "usage_events"
			format      = "csv"
			compression = "gzip"
			status      = "paused"
			time_zone   = "UTC"

			labels = {
				team = "finops"
			}
		}
	`
)

func TestAICostExportSchedule(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()

	newConfig := func(t *testing.T, resources ...string) string {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		for _, content := range resources {
			resource, err := hclbuilder.FromString(content)
			require.NoError(t, err)
			builder.Upsert(resource)
		}

		return builder.Build()
	}

	t.Run("Daily Usage Schedule", func(t *testing.T) {
		const address = "konnect_ai_cost_export_schedule.test_schedule"

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: newConfig(t, aiCostExportSchedule),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(address, "id"),
						resource.TestCheckResourceAttr(address, "name", "tf-test-daily-usage"),
						resource.TestCheckResourceAttr(address, "cadence", "daily"),
						resource.TestCheckResourceAttr(address, "dataset", "usage_events"),
						resource.TestCheckResourceAttr(address, "format", "csv"),
						resource.TestCheckResourceAttr(address, "compression", "gzip"),
						resource.TestCheckResourceAttr(address, "status", "active"),
						resource.TestCheckResourceAttr(address, "labels.team", "finops"),
						resource.TestCheckResourceAttrSet(address, "next_run_at"),
					),
				},
				{
					Config: newConfig(t, aiCostExportSchedule),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					// A schedule is replaced, not recreated: neither `name` nor `status`
					// requires replacement.
					Config: newConfig(t, aiCostExportScheduleUpdated),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(address, "name", "tf-test-daily-usage-paused"),
						resource.TestCheckResourceAttr(address, "status", "paused"),
					),
				},
				{
					// `cadence` and `time_zone` are immutable once the schedule exists,
					// so this has to destroy and create rather than update.
					Config: newConfig(t, aiCostExportScheduleWeekly),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionReplace),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(address, "cadence", "weekly"),
					),
				},
			},
		})
	})
}
