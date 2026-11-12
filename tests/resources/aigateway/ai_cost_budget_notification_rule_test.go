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
	aiCostRuleBudget = `
		resource "konnect_ai_cost_budget" "test_rule_budget" {
			limit = {
				name = "tf-test-rule-budget"

				currency = "USD"
				value    = "1000"

				window = {
					interval = "P1M"
					anchored_at = "2026-01-01T00:00:00+01:00[Europe/Budapest]"
				}

				scope = {
					filters = jsonencode([
						{ field = "provider", op = "eq", value = "openai" },
					])
				}
			}
		}
	`

	aiCostRuleChannel = `
		resource "konnect_ai_cost_notification_channel" "test_rule_channel" {
			email = {
				name = "tf-test-rule-channel"

				recipients = ["finops@example.com"]
			}
		}
	`

	// A second channel created in the same plan. Both ids are unknown at plan time,
	// so the rule's channel list has to tolerate two unknown entries.
	aiCostRuleSlackChannel = `
		resource "konnect_ai_cost_notification_channel" "test_rule_slack_channel" {
			slack = {
				name = "tf-test-rule-slack-channel"

				webhook_url = "https://hooks.slack.com/services/T0000000/B0000000/XXXXXXXXXXXXXXXXXXXXXXXX"
			}
		}
	`

	aiCostBudgetNotificationRule = `
		resource "konnect_ai_cost_budget_notification_rule" "test_rule" {
			budget_id = konnect_ai_cost_budget.test_rule_budget.id

			channels = [
				{
					id = konnect_ai_cost_notification_channel.test_rule_channel.id
				},
				{
					id = konnect_ai_cost_notification_channel.test_rule_slack_channel.id
				}
			]

			thresholds = [
				{ percent = 50 },
				{ percent = 80 },
			]

			message = {
				subject = "{{budget.name}} is at {{threshold}}"
				body    = "Spend is {{percent}} of {{limit}}."
			}
		}
	`

	aiCostBudgetNotificationRuleUpdated = `
		resource "konnect_ai_cost_budget_notification_rule" "test_rule" {
			budget_id = konnect_ai_cost_budget.test_rule_budget.id

			channels = [
				{
					id = konnect_ai_cost_notification_channel.test_rule_channel.id
				},
				{
					id = konnect_ai_cost_notification_channel.test_rule_slack_channel.id
				}
			]

			thresholds = [
				{ percent = 60 },
				{ percent = 90 },
				{ percent = 100 },
			]

			message = {
				subject = "{{budget.name}} is at {{threshold}}"
				body    = "Spend is {{percent}} of {{limit}}."
			}
		}
	`
	// The same channels and thresholds as the base config, in the opposite order.
	// The API sorts both on read, so a set has to plan this as no change.
	aiCostBudgetNotificationRuleReordered = `
		resource "konnect_ai_cost_budget_notification_rule" "test_rule" {
			budget_id = konnect_ai_cost_budget.test_rule_budget.id

			channels = [
				{
					id = konnect_ai_cost_notification_channel.test_rule_slack_channel.id
				},
				{
					id = konnect_ai_cost_notification_channel.test_rule_channel.id
				}
			]

			thresholds = [
				{ percent = 80 },
				{ percent = 50 },
			]

			message = {
				subject = "{{budget.name}} is at {{threshold}}"
				body    = "Spend is {{percent}} of {{limit}}."
			}
		}
	`
)

func TestAICostBudgetNotificationRule(t *testing.T) {
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

	t.Run("Budget Notification Rule", func(t *testing.T) {
		const address = "konnect_ai_cost_budget_notification_rule.test_rule"

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: newConfig(t, aiCostRuleBudget, aiCostRuleChannel, aiCostRuleSlackChannel, aiCostBudgetNotificationRule),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(address, "budget_id"),
						resource.TestCheckResourceAttr(address, "channels.#", "2"),
						resource.TestCheckResourceAttr(address, "thresholds.#", "2"),
						resource.TestCheckTypeSetElemNestedAttrs(address, "thresholds.*", map[string]string{
							"percent": "50",
						}),
						resource.TestCheckResourceAttr(address, "message.subject", "{{budget.name}} is at {{threshold}}"),
					),
				},
				{
					// The rule is a singleton under its budget, so a second apply must
					// adopt the existing rule instead of planning a change.
					Config: newConfig(t, aiCostRuleBudget, aiCostRuleChannel, aiCostRuleSlackChannel, aiCostBudgetNotificationRule),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					// Same values in a different order must not diff: the API returns
					// channels sorted by name and thresholds sorted by percent.
					Config: newConfig(t, aiCostRuleBudget, aiCostRuleChannel, aiCostRuleSlackChannel, aiCostBudgetNotificationRuleReordered),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: newConfig(t, aiCostRuleBudget, aiCostRuleChannel, aiCostRuleSlackChannel, aiCostBudgetNotificationRuleUpdated),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(address, "thresholds.#", "3"),
						resource.TestCheckTypeSetElemNestedAttrs(address, "thresholds.*", map[string]string{
							"percent": "60",
						}),
						resource.TestCheckTypeSetElemNestedAttrs(address, "thresholds.*", map[string]string{
							"percent": "100",
						}),
					),
				},
			},
		})
	})
}
