resource "konnect_portal" "test_portal" {
  name                      = "tftest_portal_for_forms"
  auto_approve_developers   = false
  auto_approve_applications = false
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
