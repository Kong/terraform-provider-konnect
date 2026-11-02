package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

const (
	aiCostLimitBudget = `
		resource "konnect_ai_cost_budget" "test_limit_budget" {
			limit = {
				name     = "tf-test-limit-budget"
				posture  = "advisory"
				priority = 50

				limit = {
					currency = "USD"
					value    = "5000"
				}

				window = {
					interval = "P1M"
					anchor   = "2026-01-01T00:00:00+01:00[Europe/Budapest]"
				}

				scope = {
					each = "model"
					filters = jsonencode([
						{ field = "provider", op = "eq", value = "openai" },
					])
				}
			}
		}
	`

	// The same budget after an update. `scope` and `window` can't change on a
	// replacement, so every step repeats them exactly.
	aiCostLimitBudgetUpdated = `
		resource "konnect_ai_cost_budget" "test_limit_budget" {
			limit = {
				name     = "tf-test-limit-budget-updated"
				posture  = "enforce"
				priority = 25

				limit = {
					currency = "USD"
					value    = "7500"
				}

				window = {
					interval = "P1M"
					anchor   = "2026-01-01T00:00:00+01:00[Europe/Budapest]"
				}

				scope = {
					each = "model"
					filters = jsonencode([
						{ field = "provider", op = "eq", value = "openai" },
					])
				}
			}
		}
	`

	// A change to `window` forces a replacement: the API refuses it on an existing
	// budget. Without the `x-speakeasy-param-force-new` annotation the provider would
	// send an update and the API would answer 400.
	aiCostLimitBudgetRewindowed = `
		resource "konnect_ai_cost_budget" "test_limit_budget" {
			limit = {
				name     = "tf-test-limit-budget-daily"
				posture  = "enforce"
				priority = 25

				limit = {
					currency = "USD"
					value    = "7500"
				}

				window = {
					interval = "P1D"
					anchor   = "2026-01-01T00:00:00+01:00[Europe/Budapest]"
				}

				scope = {
					each = "model"
					filters = jsonencode([
						{ field = "provider", op = "eq", value = "openai" },
					])
				}
			}
		}
	`

	aiCostExemptBudget = `
		resource "konnect_ai_cost_budget" "test_exempt_budget" {
			exempt = {
				name     = "tf-test-exempt-budget"
				priority = 30

				scope = {
					filters = jsonencode([
						{ field = "tags.tier", op = "eq", value = "free" },
					])
				}

				validity = {
					from  = "2026-01-01T00:00:00Z"
					until = "2036-01-01T00:00:00Z"
				}
			}
		}
	`

	aiCostExemptBudgetUpdated = `
		resource "konnect_ai_cost_budget" "test_exempt_budget" {
			exempt = {
				name     = "tf-test-exempt-budget-updated"
				priority = 60

				scope = {
					filters = jsonencode([
						{ field = "tags.tier", op = "eq", value = "free" },
					])
				}

				validity = {
					from  = "2026-01-01T00:00:00Z"
					until = "2037-01-01T00:00:00Z"
				}
			}
		}
	`
)

func TestAICostBudget(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()

	// A new parent builder per step: the resource blocks are shared by pointer,
	// so reusing one parent would duplicate them.
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

	t.Run("Limit Budget", func(t *testing.T) {
		const address = "konnect_ai_cost_budget.test_limit_budget"

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: newConfig(t, aiCostLimitBudget),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(address, "id"),
						resource.TestCheckResourceAttr(address, "name", "tf-test-limit-budget"),
						resource.TestCheckResourceAttr(address, "limit.name", "tf-test-limit-budget"),
						resource.TestCheckResourceAttr(address, "limit.posture", "advisory"),
						resource.TestCheckResourceAttr(address, "limit.priority", "50"),
						resource.TestCheckResourceAttr(address, "limit.limit.currency", "USD"),
						resource.TestCheckResourceAttr(address, "limit.limit.value", "5000"),
						resource.TestCheckResourceAttr(address, "limit.scope.each", "model"),
						resource.TestCheckResourceAttr(address, "limit.window.interval", "P1M"),
					),
				},
				{
					// The API echoes the stored budget back, so an identical config must
					// produce an empty plan. A diff here means the schema does not converge.
					Config: newConfig(t, aiCostLimitBudget),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: newConfig(t, aiCostLimitBudgetUpdated),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						// The hoisted top-level `name`, `posture` and `priority` mirrors keep
						// their previous value on an update (the plan modifier prefers the
						// state value), so the assertions use the nested attributes.
						resource.TestCheckResourceAttr(address, "limit.name", "tf-test-limit-budget-updated"),
						resource.TestCheckResourceAttr(address, "limit.posture", "enforce"),
						resource.TestCheckResourceAttr(address, "limit.priority", "25"),
						resource.TestCheckResourceAttr(address, "limit.limit.value", "7500"),
					),
				},
				{
					// `window` is immutable once the budget exists, so this has to
					// destroy and create rather than update.
					Config: newConfig(t, aiCostLimitBudgetRewindowed),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionReplace),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(address, "limit.window.interval", "P1D"),
						resource.TestCheckResourceAttr(address, "limit.name", "tf-test-limit-budget-daily"),
					),
				},
			},
		})
	})

	t.Run("Exempt Budget", func(t *testing.T) {
		const address = "konnect_ai_cost_budget.test_exempt_budget"

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: newConfig(t, aiCostExemptBudget),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(address, "id"),
						resource.TestCheckResourceAttr(address, "name", "tf-test-exempt-budget"),
						resource.TestCheckResourceAttr(address, "exempt.name", "tf-test-exempt-budget"),
						resource.TestCheckResourceAttr(address, "exempt.priority", "30"),
						resource.TestCheckResourceAttr(address, "exempt.validity.from", "2026-01-01T00:00:00Z"),
						resource.TestCheckResourceAttr(address, "exempt.validity.until", "2036-01-01T00:00:00Z"),
					),
				},
				{
					Config: newConfig(t, aiCostExemptBudget),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: newConfig(t, aiCostExemptBudgetUpdated),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						// The top-level `name`, `priority` and `posture` are computed mirrors
						// hoisted from the variant block. They keep the values from creation and
						// do not refresh on update, so assert the nested attributes instead.
						resource.TestCheckResourceAttr(address, "exempt.name", "tf-test-exempt-budget-updated"),
						resource.TestCheckResourceAttr(address, "exempt.priority", "60"),
						resource.TestCheckResourceAttr(address, "exempt.validity.until", "2037-01-01T00:00:00Z"),
					),
				},
			},
		})
	})
}
