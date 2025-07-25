
resource "konnect_catalog_service" "my_catalogservice" {
  name         = "terraform-test-service"
  display_name = "Terraform Test Service"
  description  = "Hello World"
  custom_fields = jsonencode({
    "cost_center"     = "eng",
    "jira_project"    = null,
    "owner"           = "owner-name"
    "product_manager" = "pm-name"
    "slack_channel"   = null
    "dashboard" : {
      "name" : "Example dashboard",
      "link" : "https://app.example.com/dashboard/123",
    },
    "git_repo" : {
      "name" : "example-repo",
      "link" : "https://github.com/example/repo",
    },
  })

  labels = {
    key = "value"
  }
}