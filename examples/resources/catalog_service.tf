resource "konnect_catalog_service" "my_catalogservice" {
  name         = "tf-service"
  display_name = "Terraform Service"
  description  = "Hello World"
  custom_fields = jsonencode({
    "cost_center" : "eng",
    "owner" : "MEE",
    "product_manager" : "Your Name",
    "dashboard" : {
      "name" : "Example Dashboard",
      "link" : "https://app.example.com/dashboard/123",
    },
    "git_repo" : {
      "name" : "example-repo",
      "link" : "https://github.com/example/repo",
    },
    "jira_project" : null,
    "slack_channel" : {
      "name" : "test-channel",
      "link" : "https://example.slack.com/archives/C098WKDB020"
    },
  })

  labels = {
    key = "value"
  }
}
