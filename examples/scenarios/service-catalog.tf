# Don't forget to create auth.tf to configure the provider
# (see examples/scenarios/_auth.tf for an example)

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

resource "konnect_integration_instance" "gitlab" {
  name         = "my-gitlab"
  display_name = "My GitLab"
  description  = "GitLab cloud"

  integration_name = "gitlab"
  config = jsonencode({
    base_url = "https://gitlab.com/api/v4"
  })
}

resource "konnect_integration_instance_auth_config" "my_integrationinstanceauthconfig" {
  integration_instance_id = konnect_integration_instance.gitlab.id
  oauth_config = {
    authorization_endpoint = "https://identity.example.com/oauth/authorize"
    client_id              = "d745213a-b7e8-4998-abe3-41f164001970"
    client_secret          = "s3cr3t4p1cl13ntt0k3n1234567890abcdef2"
    token_endpoint         = "https://identity.example.com/oauth/token"
  }
}

resource "konnect_integration_instance" "swagger" {
  name         = "my-swaggerhub"
  display_name = "My SwaggerHub"
  description  = "Swaggerhub account which hosts internal APIs only."

  integration_name = "swaggerhub"
  config           = jsonencode({})
}

resource "konnect_integration_instance_auth_credential" "my_integrationinstanceauthcredential" {
  integration_instance_id = konnect_integration_instance.swagger.id
  multi_key_auth = {
    config = {
      headers = [
        {
          name = "authorization"
          key  = "swaggerhub-api-key"
        }
      ]
    }
  }
}

