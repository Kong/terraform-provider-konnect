resource "konnect_application_auth_strategy" "test_auth_strategy" {
  key_auth = {
    name          = "auth_strategy_test_with_id"
    display_name  = "Test Key Auth With ID"
    strategy_type = "key_auth"
    configs = {
      key_auth = {
        key_names = ["apikey"]
      }
    }
  }
}

resource "konnect_portal" "test_portal_app" {
  name                                 = "test_portal_for_app_auth_id"
  default_api_visibility               = "public"
  auto_approve_developers              = true
  auto_approve_applications            = true
}

resource "konnect_api" "test_api_auth" {
  name    = "Test API Auth Strategy"
  slug    = "test-api-auth-strategy-v1"
  version = "v1"
}

resource "konnect_api_publication" "test_pub_auth" {
  api_id                     = konnect_api.test_api_auth.id
  portal_id                  = konnect_portal.test_portal_app.id
  visibility                 = "public"
  auto_approve_registrations = true
  auth_strategy_ids          = [konnect_application_auth_strategy.test_auth_strategy.id]
}

resource "konnect_portal_developer" "test_app_developer" {
  portal_id             = konnect_portal.test_portal_app.id
  email                 = "app.developer.auth@example.com"
  full_name             = "Application Developer Auth"
  status                = "approved"
  send_invitation_email = false
}

resource "konnect_portal_application" "test_application" {
  portal_id        = konnect_portal.test_portal_app.id
  auth_strategy_id = konnect_application_auth_strategy.test_auth_strategy.id
  name             = "test_application_auth_id"
  description      = "Test application with explicit auth strategy"

  owner = {
    id   = konnect_portal_developer.test_app_developer.id
    type = "developer"
  }
}

