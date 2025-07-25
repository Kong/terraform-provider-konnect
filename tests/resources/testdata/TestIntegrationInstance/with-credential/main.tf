resource "konnect_integration_instance" "my_integration_instance" {
  name         = "my-swaggerhub"
  display_name = "My SwaggerHub"
  description  = "Some Description"

  integration_name = "swaggerhub"
  config           = jsonencode({})
}

resource "konnect_integration_instance_auth_credential" "my_integrationinstanceauthcredential" {
  integration_instance_id = konnect_integration_instance.my_integration_instance.id
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
