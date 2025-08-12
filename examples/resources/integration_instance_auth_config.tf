resource "konnect_integration_instance_auth_config" "my_integrationinstanceauthconfig" {
  integration_instance_id = konnect_integration_instance.my_integration_instance.id
  oauth_config = {
    authorization_endpoint = "https://identity.example.com/oauth/authorize"
    client_id              = "d745213a-b7e8-4998-abe3-41f164001970"
    client_secret          = "s3cr3t4p1cl13ntt0k3n1234567890abcdef2"
    token_endpoint         = "https://identity.example.com/oauth/token"
  }
}
