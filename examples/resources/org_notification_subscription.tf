resource "konnect_org_notification_subscription" "example" {
  destination_ids = [konnect_org_notification_destination.example.id]
  enabled         = true
  entities        = ["*"]
  event_type      = "kai.feature.enablement"
  name            = "example-subscription"
  regions         = ["*"]
}
