resource "konnect_org_notification_destination" "example" {
  http = {
    name = "example-webhook"
    type = "http"
    url  = "https://example.com/notifications"

    custom_headers = {
      "X-Environment" = "production"
    }
  }
}
