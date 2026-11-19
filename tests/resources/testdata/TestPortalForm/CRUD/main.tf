resource "konnect_portal" "test_portal" {
  name                      = "tftest_portal_for_forms"
  auto_approve_developers   = false
  auto_approve_applications = false
  default_application_auth_strategy_id = konnect_application_auth_strategy.test_auth_strategy.id
}

resource "konnect_portal_form" "test_form" {
  portal_id = konnect_portal.test_portal.id

  developer_registration = {
    status = "published"
    fields = [
      {
        text = {
          name  = "full_name"
          label = "nombre"
        }
      },
      {
        email = {
          name  = "email"
          label = "emaile"
        }
      },
      {
        text = {
          name     = "custom"
          label    = "custom"
          required = true
        }
      },
      {
        text = {
          name     = "custom2"
          label    = "custom two"
          required = true
        }
      },
      {
        submit = {
          name  = "submit"
          value = "x"
        }
      },
    ]
  }
}

resource "konnect_portal_developer" "test_developer" {
  portal_id             = konnect_portal.test_portal.id
  email                 = "tftestformdev@example.com"
  full_name             = "TFTest Form Developer"
  send_invitation_email = false
  status                = "approved"

  additional_data = jsonencode({
    custom  = "tftestcustomvalue"
    custom2 = "tftestcustomvalue2"
  })

  depends_on = [konnect_portal_form.test_form]
}

resource "konnect_application_auth_strategy" "test_auth_strategy" {
  key_auth = {
    name          = "tftest-key-auth-forms"
    display_name  = "TFTest Key Auth Forms"
    strategy_type = "key_auth"
    configs = {
      key_auth = {
        key_names = ["apikey"]
      }
    }
  }
}

resource "konnect_portal_form" "test_api_form" {
  portal_id = konnect_portal.test_portal.id

  api_registration = {
    name   = "tftest-api-registration"
    status = "published"
    fields = [
      {
        text = {
          name     = "api_id"
          label    = "API"
          required = true
        }
      },
      {
        textarea = {
          name     = "use_case"
          label    = "Use case"
          required = true
        }
      },
      {
        submit = {
          name  = "submit"
          value = "Request access"
        }
      },
    ]
  }
}

resource "konnect_api" "test_api" {
  name    = "TFTest Forms API"
  slug    = "tftest-forms-api"
  version = "v1"
}

resource "konnect_api_publication" "test_api_publication" {
  api_id                     = konnect_api.test_api.id
  portal_id                  = konnect_portal.test_portal.id
  visibility                 = "public"
  auto_approve_registrations = false
  auth_strategy_ids          = [konnect_application_auth_strategy.test_auth_strategy.id]
  form_id                    = konnect_portal_form.test_api_form.api_registration.id
}

resource "konnect_portal_application" "test_application" {
  portal_id   = konnect_portal.test_portal.id
  name        = "tftest_forms_application"
  description = "Application for form registration test"

  owner = {
    id   = konnect_portal_developer.test_developer.id
    type = "developer"
  }
}

resource "konnect_portal_application_registration" "test_registration" {
  portal_id      = konnect_portal.test_portal.id
  api_id         = konnect_api.test_api.id
  application_id = konnect_portal_application.test_application.id
  status         = "approved"

  additional_data = jsonencode({
    use_case = "tftest use case"
  })

  depends_on = [konnect_api_publication.test_api_publication]
}
