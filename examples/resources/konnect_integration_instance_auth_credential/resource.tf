resource "konnect_integration_instance_auth_credential" "my_integrationinstanceauthcredential" {
  integration_instance_id = "3f51fa25-310a-421d-bd1a-007f859021a3"
  multi_key_auth = {
    config = {
      headers = [
        {
          key  = "...my_key..."
          name = "...my_name..."
        }
      ]
    }
  }
}