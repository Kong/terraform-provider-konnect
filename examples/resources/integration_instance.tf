resource "konnect_integration_instance" "my_integrationinstance" {
  name         = "my-github"
  display_name = "My GitHub"
  description  = "Some Description"

  integration_name = "github"
  config = {}
}
