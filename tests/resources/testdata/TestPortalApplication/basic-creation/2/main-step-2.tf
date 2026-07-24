resource "konnect_application_auth_strategy" "my_auth_strategy_basic" {
  key_auth = {
    name          = "key-auth-basic-test"
    display_name  = "Key Auth Basic Test"
    strategy_type = "key_auth"
    configs = {
      key_auth = {
        key_names = ["apikey"]
      }
    }
  }
}

resource "konnect_portal" "test_portal_app" {
  name                                 = "test_portal_for_app_basic"
  default_api_visibility               = "public"
  auto_approve_developers              = true
  auto_approve_applications            = true
  default_application_auth_strategy_id = konnect_application_auth_strategy.my_auth_strategy_basic.id
}

resource "konnect_api" "my_api_basic" {
  name    = "Test API Basic Creation"
  slug    = "test-api-basic-creation-v1"
  version = "v1"
}

resource "konnect_api_publication" "my_api_publication_basic" {
  api_id                     = konnect_api.my_api_basic.id
  portal_id                  = konnect_portal.test_portal_app.id
  visibility                 = "public"
  auto_approve_registrations = true
  auth_strategy_ids          = [konnect_application_auth_strategy.my_auth_strategy_basic.id]
}

resource "konnect_portal_developer" "test_app_developer" {
  portal_id             = konnect_portal.test_portal_app.id
  email                 = "app.developer.basic@example.com"
  full_name             = "Application Developer Basic"
  status                = "approved"
  send_invitation_email = false
}

resource "konnect_portal_application" "test_application" {
  portal_id   = konnect_portal.test_portal_app.id
  name        = "test_application_basic"
  description = "Basic test application"

  owner = {
    id   = konnect_portal_developer.test_app_developer.id
    type = "developer"
  }
}

