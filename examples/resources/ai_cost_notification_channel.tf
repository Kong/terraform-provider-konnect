# Notification channels are the destinations that budget alerts go to. A channel
# is either an email channel or a Slack channel, never both.
resource "konnect_ai_cost_notification_channel" "email" {
  email = {
    name = "finops-email"

    recipients = [
      "finops@example.com",
      "platform-leads@example.com",
    ]

    labels = {
      team = "finops"
    }
  }
}

resource "konnect_ai_cost_notification_channel" "slack" {
  slack = {
    name = "finops-slack"

    # Create this incoming webhook in your own Slack workspace. The API
    # returns the URL with its token replaced by `…`, so Terraform keeps the
    # configured value in state.
    webhook_url = "https://hooks.slack.com/services/T0000000/B0000000/XXXXXXXXXXXXXXXXXXXXXXXX"
  }
}
