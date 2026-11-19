# A portal with a custom developer registration form, and a developer
# that supplies values for the form's custom fields via additional_data.

resource "konnect_portal" "forms_portal" {
  name                      = "portal-with-registration-form"
  auto_approve_developers   = false
  auto_approve_applications = false
}

resource "konnect_portal_form" "registration" {
  portal_id = konnect_portal.forms_portal.id

  developer_registration = {
    status = "published"
    fields = [
      {
        text = {
          name     = "full_name"
          label    = "Full name"
          required = true
        }
      },
      {
        email = {
          name     = "email"
          label    = "Email"
          required = true
        }
      },
      {
        text = {
          name     = "company"
          label    = "Company"
          required = true
        }
      },
      {
        submit = {
          name  = "submit"
          value = "Register"
        }
      },
    ]
  }
}

# Custom field values must match the field names defined in the form above.
resource "konnect_portal_developer" "registered_developer" {
  portal_id             = konnect_portal.forms_portal.id
  email                 = "developer@example.com"
  full_name             = "Example Developer"
  send_invitation_email = false
  status                = "approved"

  additional_data = jsonencode({
    company = "Example Corp"
  })

  # The form must exist before the developer is created so that
  # additional_data is validated against it.
  depends_on = [konnect_portal_form.registration]
}
