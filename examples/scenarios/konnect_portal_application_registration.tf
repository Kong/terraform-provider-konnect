resource "konnect_application_auth_strategy" "my_auth_strategy_reg" {
  key_auth = {
    name          = "key-auth-registration-test"
    display_name  = "Key Auth Registration Test"
    strategy_type = "key_auth"
    configs = {
      key_auth = {
        key_names = ["apikey"]
      }
    }
  }
}

resource "konnect_portal" "test_portal_reg" {
  name                                 = "test_portal_for_registration"
  default_api_visibility               = "public"
  auto_approve_developers              = true
  auto_approve_applications            = true
  default_application_auth_strategy_id = konnect_application_auth_strategy.my_auth_strategy_reg.id
}

resource "konnect_api" "my_api_reg" {
  name    = "Test API Registration"
  slug    = "test-api-registration-v3-crud"
  version = "v3"
}

resource "konnect_api_publication" "my_api_publication_reg" {
  api_id                     = konnect_api.my_api_reg.id
  portal_id                  = konnect_portal.test_portal_reg.id
  visibility                 = "public"
  auto_approve_registrations = false
  auth_strategy_ids          = [konnect_application_auth_strategy.my_auth_strategy_reg.id]
}

resource "konnect_portal_developer" "test_reg_developer" {
  portal_id             = konnect_portal.test_portal_reg.id
  email                 = "reg.developer@example.com"
  full_name             = "Registration Developer"
  status                = "approved"
  send_invitation_email = false
}

resource "konnect_portal_application" "test_application_reg" {
  portal_id   = konnect_portal.test_portal_reg.id
  name        = "test_application_for_registration"
  description = "Test application for registration"

  owner = {
    id   = konnect_portal_developer.test_reg_developer.id
    type = "developer"
  }

}

resource "konnect_portal_application_registration" "test_registration" {
  portal_id = konnect_portal.test_portal_reg.id
  api_id    = konnect_api.my_api_reg.id
  status    = "approved"
  application = {
    id = konnect_portal_application.test_application_reg.id
  }
}
