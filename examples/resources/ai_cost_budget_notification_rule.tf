# A notification rule alerts channels when spend in a limit budget reaches a
# share of the limit. A budget has at most one rule, so this resource is keyed
# by the budget.
#
# Import with the budget id: terraform import konnect_ai_cost_budget_notification_rule.alerts <budget-id>
resource "konnect_ai_cost_budget_notification_rule" "alerts" {
  budget_id = konnect_ai_cost_budget.team_models.id

  channels = [
    { id = konnect_ai_cost_notification_channel.email.id },
    { id = konnect_ai_cost_notification_channel.slack.id },
  ]

  thresholds = [
    { percent = 50 },
    { percent = 80 },
    { percent = 100 },
  ]

  # The wording is optional. Placeholders are filled in for each alert.
  message = {
    subject = "{{budget.name}} is at {{threshold}}"
    body    = "Spend reached {{threshold}} of the {{limit}} limit. Current spend is {{spend}}."
  }
}
